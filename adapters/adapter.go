package adapters

import (
	"github.com/vishnusunil243/ProductService/entities"
	"gorm.io/gorm"
)

type ProductAdapter struct {
	DB *gorm.DB
}

func NewProductAdapter(db *gorm.DB) *ProductAdapter {
	return &ProductAdapter{
		DB: db,
	}
}
func (product *ProductAdapter) AddProduct(req entities.Product) (entities.Product, error) {
	var res entities.Product
	query := `INSERT INTO products(name,price,quantity) VALUES($1,$2,$3) RETURNING id,name,price,quantity`
	return res, product.DB.Raw(query, req.Name, req.Price, req.Quantity).Scan(&res).Error
}