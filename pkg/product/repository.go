package product

type (
	// Read manager read actions dataase
	Read interface {
		FindAll(filter map[string]interface{}) ([]*Product, error)
		Find(id string) (*Product, error)
	}
	// Write manager write actions database
	Write interface {
		Create(product *Product) error
		Update(product *Product) error
		Delete(product *Product) error
	}
	// Repository factory from manipulate data in database
	Repository interface {
		Read
		Write
	}
)
