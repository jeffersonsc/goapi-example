package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jeffersonsc/natureapi/pkg/product"
)

// ProductHandler setup routes from products (controllers or actions)
type ProductHandler struct {
	repo   product.Repository
	log    *log.Logger
	Router *mux.Router
}

// NewProductHandler create a new instance of handler
func NewProductHandler(repo product.Repository, router *mux.Router) ProductHandler {
	logger := NewProductHandlerLogger()
	productHandler := ProductHandler{
		repo:   repo,
		log:    logger,
		Router: router,
	}

	router.HandleFunc("/v1/products", productHandler.Index).Methods(http.MethodGet)
	router.HandleFunc("/v1/products", productHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/v1/products/{id}", productHandler.Show).Methods(http.MethodGet)
	router.HandleFunc("/v1/products/{id}", productHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/v1/products/{id}", productHandler.Delete).Methods(http.MethodDelete)

	return productHandler
}

// NewProductHandlerLogger instace a new logger
func NewProductHandlerLogger() *log.Logger {
	return log.New(os.Stdout, "[product-handler]", 0)
}

// Index method GET /products
func (ph ProductHandler) Index(w http.ResponseWriter, r *http.Request) {
	service := product.NewService(ph.repo)

	products, err := service.FindAll()
	if err != nil {
		ph.log.Println("Error on find all products ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Failed get all products, please contact suport"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Create method POST /products
func (ph ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := product.DTO{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		ph.log.Println("Failed decode json input ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	service := product.NewService(ph.repo)

	err = service.IsValid(&body)
	if err != nil {
		// Return 422
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
		return
	}

	result, err := service.Create(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Failed create a new product, please contact support"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// Show method GET /products/:id
func (ph ProductHandler) Show(w http.ResponseWriter, r *http.Request) {
	products := map[string]interface{}{"name": "test"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Update method PUT /products/:id
func (ph ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	products := map[string]interface{}{"name": "test"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Delete method DELETE /products/:id
func (ph ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	products := map[string]interface{}{"name": "test"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
