package trendyol

import (
	"bytes"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/exp/errors/fmt"
)

const (
	allBrandsURL     = "https://api.trendyol.com/sapigw/brands"
	allCategoriesURL = "https://api.trendyol.com/sapigw/product-categories"
	allProvidersURL  = "https://api.trendyol.com/sapigw/shipment-providers"
)

/*CreateProduct creates a product*/
func CreateProduct(supplierID string, productInfo []byte) {
	url := "https://api.trendyol.com/sapigw/suppliers/" +
		supplierID + "/v2/products"

	resp, err := http.Post(url, "applciation/json",
		bytes.NewBuffer(productInfo))
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(body)
}

/*UpdateProduct updates product infos*/
func UpdateProduct(supplierID string, productInfo []byte) {
	url := "https://api.trendyol.com/sapigw/suppliers/" +
		supplierID + "/v2/products"

	resp, err := http.NewRequest(http.MethodPut, url,
		bytes.NewBuffer(productInfo))
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(body)

}

/*UpdateProductPriceAndInventory updates product's prica and inventory*/
func UpdateProductPriceAndInventory(supplierID string,
	priceandinventoryInfo []byte) {
	url := "https://api.trendyol.com/sapigw/suppliers/" +
		supplierID + "/products/price-and-inventory"

	resp, err := http.Post(url, "applciation/json",
		bytes.NewBuffer(priceandinventoryInfo))
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(body)
}

/*GetAllBrands gets all brands from trendyol api*/
func GetAllBrands() (brands Brands) {
	body := makeGetRequest(allBrandsURL)

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &brands)
	return brands
}

/*GetAllCategories gets all categories from trendyol api*/
func GetAllCategories() (categories AllCategories) {
	body := makeGetRequest(allCategoriesURL)

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &categories)
	return categories
}

/*GetCategoryAttributes gets category attributes by categoryID from trendyol api*/
func GetCategoryAttributes(categoryID string) (attributes CategoryAttributes) {
	body := makeGetRequest("https://api.trendyol.com/sapigw/product-categories/" +
		categoryID + "/attributes")

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &attributes)
	return attributes
}

/*GetProviders gets providers from trendyol api*/
func GetProviders() (providers []Provider) {
	body := makeGetRequest(allProvidersURL)

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &providers)
	return providers
}

func makeGetRequest(url string) []byte {
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
