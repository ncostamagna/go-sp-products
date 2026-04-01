package product

import "errors"

var (
	ErrNameRequired = errors.New("name is required")
	ErrPriceRequired = errors.New("price is required")
	ErrPriceNegative = errors.New("price must be positive or zero")

	ErrProductNotFound = errors.New("product not found in database")

	ErrIdRequired = errors.New("id is required")
)