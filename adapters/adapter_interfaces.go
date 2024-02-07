package adapters

import "github.com/vishnusunil243/ProductService/entities"

type AdapterInterface interface {
	AddProduct(req entities.Product) (entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
}
