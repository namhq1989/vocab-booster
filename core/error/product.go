package apperrors

import "errors"

var Product = struct {
	ProductNotFound      error
	InvalidProduct       error
	InvalidPublisher     error
	InvalidSupplier      error
	InvalidBrand         error
	InvalidCategory      error
	InvalidMinPrice      error
	InvalidMaxPrice      error
	InvalidListedPrice   error
	InvalidStimulusPrice error
	InvalidPrice         error
	InvalidCommission    error
	InvalidSKU           error
	InvalidSKUPrice      error
}{
	ProductNotFound:      errors.New("product_not_found"),
	InvalidProduct:       errors.New("product_invalid_product"),
	InvalidPublisher:     errors.New("product_invalid_publisher"),
	InvalidSupplier:      errors.New("product_invalid_supplier"),
	InvalidBrand:         errors.New("product_invalid_brand"),
	InvalidCategory:      errors.New("product_invalid_category"),
	InvalidMinPrice:      errors.New("product_invalid_min_price"),
	InvalidMaxPrice:      errors.New("product_invalid_max_price"),
	InvalidListedPrice:   errors.New("product_invalid_listed_price"),
	InvalidStimulusPrice: errors.New("product_invalid_stimulus_price"),
	InvalidCommission:    errors.New("product_invalid_commission"),
	InvalidPrice:         errors.New("product_invalid_price"),
	InvalidSKU:           errors.New("product_invalid_sku"),
	InvalidSKUPrice:      errors.New("product_invalid_sku_price"),
}
