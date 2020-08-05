package trendyol

import (
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/exp/errors/fmt"
)

const (
	allBrandsURL     = "https://api.trendyol.com/sapigw/brands"
	allCategoriesURL = "https://api.trendyol.com/sapigw/product-categories"
)

// Brands stores Trendyol brands infos(trendyol response)
type Brands struct {
	Brands []Brand `json:"brands"`
}

// Brand stores a Trendyol brand infos(trendyol response)
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// AllCategories is a root for categories
type AllCategories struct {
	Categories []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SubCategories []struct {
			ID            int           `json:"id"`
			Name          string        `json:"name"`
			ParentID      int           `json:"parentId"`
			SubCategories []interface{} `json:"subCategories"`
		} `json:"subCategories"`
	} `json:"categories"`
}

// CategoryAttributes stores a Trendyol Category Attributes infos
type CategoryAttributes struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	DisplayName        string `json:"displayName"`
	CategoryAttributes []struct {
		CategoryID int `json:"categoryId"`
		Attribute  struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"attribute"`
		Required        bool `json:"required"`
		AllowCustom     bool `json:"allowCustom"`
		Varianter       bool `json:"varianter"`
		Slicer          bool `json:"slicer"`
		AttributeValues []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"attributeValues"`
	} `json:"categoryAttributes"`
}

/*GetAllBrands gets all brands from trendyol api*/
func GetAllBrands() (brands Brands) {
	body := makeRequest(allBrandsURL)

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &brands)
	return brands
}

/*GetAllCategories gets all categories from trendyol api*/
func GetAllCategories() (categories AllCategories) {
	body := makeRequest(allCategoriesURL)

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &categories)
	return categories
}

/*GetCategoryAttributes gets category attributes by categoryID from trendyol api*/
func GetCategoryAttributes(categoryID string) (attributes CategoryAttributes) {
	body := makeRequest("https://api.trendyol.com/sapigw/product-categories/" +
		categoryID + "/attributes")

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &attributes)
	fmt.Println(attributes)
	return attributes
}

func makeRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	return body
}
