package product

// DTO comunicate service and controller application
type DTO struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty" validate:"required"`
	Description  string   `json:"description,omitempty" validate:"required"`
	Price        float64  `json:"price,omitempty" validate:"required,gte=0,numeric"`
	Images       []string `json:"images,omitempty" validate:"required"`
	CurrencyCode string   `json:"currency_code,omitempty" validate:"required,gte=3,lte=3,alpha"`
}

// ToProduct transtform DTO to Product
func (dto *DTO) ToProduct() *Product {
	return &Product{
		Name:         dto.Name,
		Description:  dto.Description,
		CurrencyCode: dto.CurrencyCode,
		Images:       dto.Images,
		Price:        dto.Price,
	}
}

// FromProduct transform Produtc to DTO
func FromProduct(product *Product) *DTO {
	return &DTO{
		ID:           product.ID.Hex(),
		Name:         product.Name,
		Description:  product.Description,
		CurrencyCode: product.CurrencyCode,
		Images:       product.Images,
		Price:        product.Price,
	}
}
