package adapters

import "github.com/vishnusunil243/ProductService/entities"

type AdapterInterface interface {
	AddProduct(req entities.Product) (entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
	GetProduct(id uint) (entities.Product, error)
	IncrementQuantity(id uint, quantity int) (entities.Product, error)
	DecrementQuantity(id uint, quantity int) (entities.Product, error)
}
