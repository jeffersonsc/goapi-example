package product

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
)

// Service layer intermediate data transit
type Service struct {
	repo      Repository
	validator *validator.Validate
	log       *log.Logger
}

// ErrValidationError return case is not possible validate input data
var (
	ErrValidationError = fmt.Errorf("Sorry it not possible validate input data")
	ErrCreateProduct   = fmt.Errorf("Sorry failed create a new product")
	ErrUpdateProduct   = fmt.Errorf("Sorry failed update product")
	ErrDeleteProduct   = fmt.Errorf("Sorry failed delete product")
	ErrFindProduct     = fmt.Errorf("Sorry failed find a product especificed")
)

// NewService instance a new service with repository
func NewService(repo Repository) *Service {
	logger := NewServiceLogger()
	return &Service{
		repo: repo,
		log:  logger,
	}
}

// NewServiceLogger generate new logger from server
func NewServiceLogger() *log.Logger {
	return log.New(os.Stdout, "[product.service]", 0)
}

// IsValid check playload as valid
func (s Service) IsValid(dto *DTO) error {
	// Validate input data
	err := s.validator.Struct(dto)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ErrValidationError
		}

		return err
	}

	return nil
}

// FindAll products and transform to DTO
func (s Service) FindAll() ([]*DTO, error) {
	products, err := s.repo.FindAll(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	result := []*DTO{}
	for _, product := range products {
		result = append(result, FromProduct(product))
	}

	return result, nil
}

// Find product by id
func (s Service) Find(id string) (*DTO, error) {
	product, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}

	return FromProduct(product), nil
}

// Create a new product and revalidate cache
func (s Service) Create(dto *DTO) (*DTO, error) {
	product := dto.ToProduct()

	err := s.repo.Create(product)
	if err != nil {
		s.log.Println("[Service - Create] Failed create a new prduct ERROR:", err.Error())
		return nil, ErrCreateProduct
	}

	return FromProduct(product), nil
}

// Update product and revalidate cache
func (s Service) Update(dto *DTO) (*DTO, error) {
	product := dto.ToProduct()

	err := s.repo.Update(product)
	if err != nil {
		s.log.Println("[Service - Update] Failed update prduct ERROR:", err.Error())
		return nil, ErrUpdateProduct
	}

	return FromProduct(product), nil
}

// Delete product and revalidate cache
func (s Service) Delete(dto *DTO) (*DTO, error) {
	product := dto.ToProduct()

	err := s.repo.Update(product)
	if err != nil {
		s.log.Println("[Service - Delete] Failed delete prduct ERROR:", err.Error())
		return nil, ErrDeleteProduct
	}

	return FromProduct(product), nil
}
