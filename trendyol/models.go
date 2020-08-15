package trendyol

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

// Provider stores a Trendyol Provider infos
type Provider struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	TaxNumber string `json:"taxNumber"`
}

// Product used when create or update product
type Product struct {
	Items []struct {
		Barcode           string  `json:"barcode"`
		Title             string  `json:"title"`
		ProductMainID     string  `json:"productMainId"`
		BrandID           int     `json:"brandId"`
		CategoryID        int     `json:"categoryId"`
		Quantity          int     `json:"quantity"`
		StockCode         string  `json:"stockCode"`
		DimensionalWeight int     `json:"dimensionalWeight"`
		Description       string  `json:"description"`
		CurrencyType      string  `json:"currencyType"`
		ListPrice         float64 `json:"listPrice"`
		SalePrice         float64 `json:"salePrice"`
		VatRate           int     `json:"vatRate"`
		CargoCompanyID    int     `json:"cargoCompanyId"`
		Images            []struct {
			URL string `json:"url"`
		} `json:"images"`
		Attributes []struct {
			AttributeID          int    `json:"attributeId"`
			AttributeValueID     int    `json:"attributeValueId,omitempty"`
			CustomAttributeValue string `json:"customAttributeValue,omitempty"`
		} `json:"attributes"`
	} `json:"items"`
}

// PriceAndInventory used when update price and inventory
type PriceAndInventory struct {
	Items []struct {
		Barcode   string  `json:"barcode"`
		Quantity  int     `json:"quantity"`
		SalePrice float64 `json:"salePrice"`
		ListPrice float64 `json:"listPrice"`
	} `json:"items"`
}

type Response struct {
	BatchRequestID string `json:"batchRequestId"`
}

type BatchRequestResult struct {
	BatchRequestID string `json:"batchRequestId"`
	Items          []struct {
		RequestItem struct {
			Product struct {
				Brand        string  `json:"brand"`
				Barcode      string  `json:"barcode"`
				Title        string  `json:"title"`
				Description  string  `json:"description"`
				CategoryName string  `json:"categoryName"`
				ListPrice    float64 `json:"listPrice"`
				SalePrice    float64 `json:"salePrice"`
				CurrencyType string  `json:"currencyType"`
				VatRate      int     `json:"vatRate"`
				CargoCompany string  `json:"cargoCompany"`
				Quantity     int     `json:"quantity"`
				StockCode    string  `json:"stockCode"`
				Images       []struct {
					URL string `json:"url"`
				} `json:"images"`
				ProductMainID     string        `json:"productMainId"`
				Gender            string        `json:"gender"`
				DimensionalWeight int           `json:"dimensionalWeight"`
				Attributes        []interface{} `json:"attributes"`
				VariantAttributes []struct {
					AttributeName  string `json:"attributeName"`
					AttributeValue string `json:"attributeValue"`
				} `json:"variantAttributes"`
			} `json:"product"`
		} `json:"requestItem"`
		Status         string        `json:"status"`
		FailureReasons []interface{} `json:"failureReasons"`
	} `json:"items"`
	Status           string `json:"status"`
	CreationDate     int64  `json:"creationDate"`
	LastModification int64  `json:"lastModification"`
	SourceType       string `json:"sourceType"`
}
