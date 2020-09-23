package trendyol

import (
	"bytes"
	"entegrasyon/templates"
	"io/ioutil"
	"net/http"

	"fmt"

	jsoniter "github.com/json-iterator/go"
)

const (
	allBrandsURL     = "https://api.trendyol.com/sapigw/brands"
	allCategoriesURL = "https://api.trendyol.com/sapigw/product-categories"
	allProvidersURL  = "https://api.trendyol.com/sapigw/shipment-providers"
)

func GetLayout(w http.ResponseWriter, r *http.Request) {
	err := templates.RenderInLayout(w, "pages-settings.html", nil)
	if err != nil {
		panic(err)
	}
}

/*CreateProduct creates a product*/
func CreateProduct(supplierID string, productInfo []byte) BatchRequestResult {
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

	var res Response
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &res)

	return getBatchRequestResult(supplierID, res.BatchRequestID)
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
	priceandinventoryInfo []byte) BatchRequestResult {
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

	var res BatchRequestResult
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &res)

	return getBatchRequestResult(supplierID, res.BatchRequestID)
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

func getBatchRequestResult(supplierID, batchRequestID string) (result BatchRequestResult) {
	url := "https://api.trendyol.com/sapigw/suppliers/" +
		supplierID + "/products/batch-requests/" + batchRequestID
	body := makeGetRequest(url)

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &result)
	return result
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
