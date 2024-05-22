package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/goexpert/9-APIS/internal/dto"
	"github.com/devfullcycle/goexpert/9-APIS/internal/entity"
	"github.com/devfullcycle/goexpert/9-APIS/internal/infra/database"
)

type ProductHandler struct {
	productDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		productDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.productDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
