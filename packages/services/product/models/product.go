package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type AssetType string

type JSON json.RawMessage

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to parse attributes")
	}
	result := json.RawMessage{}

	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

const (
	IMAGE AssetType = "IMAGE"
	VIDEO AssetType = "VIDEO"
	EMBED AssetType = "EMBED"
	PDF   AssetType = "PDF"
	TEXT  AssetType = "TEXT"
)

type Attribute struct{}
type PriceInfo struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

type FulfilmentFlatRates struct{}

func (p *PriceInfo) Scan(value interface{}) error {
	if raw, ok := value.(string); ok {
		result := strings.Replace(raw, "(", "", -1)
		result = strings.Replace(result, ")", "", -1)
		resultArr := strings.Split(result, ",")
		amount, err := strconv.ParseFloat(resultArr[0], 32)
		if err != nil {
			return errors.New(fmt.Sprint("failed to scan PriceInfo value:", value))
		}
		priceInfo := PriceInfo{Amount: float32(amount), Currency: resultArr[1]}
		*p = PriceInfo(priceInfo)
		return nil
	} else {
		return errors.New(fmt.Sprint("failed to scan PriceInfo value:", value))
	}

}

func (p PriceInfo) Value() (driver.Value, error) {

	value := fmt.Sprintf("(%f,%s)", p.Amount, p.Currency)
	return value, nil
}

type BasicProduct interface {
	Details() ProductDetails
	IsAvailable() bool
	Type() string
	HasMedia() bool
	GetMedia() []ProductAsset
}
type ProductDetails struct {
	Id                  string    `json:"id"`
	ActiveStartDate     time.Time `json:"active_start_date"`
	ActiveEndDate       time.Time `json:"active_end_date"`
	Attributes          JSON      `json:"attributes"`
	Cost                PriceInfo `json:"cost"`
	Currency            string    `json:"currency"`
	DefaultPrice        PriceInfo `json:"default_price"`
	LongDescription     string    `json:"long_description"`
	Discountable        bool      `json:"discountable"`
	ExternalId          string    `json:"external_id"`
	FulfilmentFlatRates JSON      `json:"fulfilment_flat_rates"`
	Keywords            []string  `json:"keywords"`
	MetaDescription     string    `json:"meta_description"`
	MetaTitle           string    `json:"meta_title"`
	SalePrice           PriceInfo `json:"sale_price"`
	ShortDescription    string    `json:"short_description"`
	SKU                 string    `json:"sku"`
	Tags                string    `json:"tags"`
	Title               string    `json:"title"`
	UPC                 string    `json:"upc"`
	URI                 string    `json:"uri"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type ProductAsset struct {
	Id      int       `json:"id"`
	Sorting string    `json:"sorting"`
	Type    AssetType `json:"type"`
	URL     string    `json:"url"`
	AltText string    `json:"alt_text"`
	Title   string    `json:"title"`
}

type ProductVariant struct {
	ActiveEndDate   time.Time              `json:"active_end_date"`
	ActiveStartDate time.Time              `json:"active_start_date"`
	Attributes      JSON                   `json:"attributes"`
	Cost            float32                `json:"cost"`
	DefaultPrice    float32                `json:"default_price"`
	Description     string                 `json:"description"`
	OptionValues    map[string]interface{} `json:"option_values"`
	SKU             string                 `json:"sku"`
	UPC             string                 `json:"upc"`
}
