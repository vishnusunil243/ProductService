package adapters

import (
	"fmt"

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
func (product *ProductAdapter) GetAllProducts() ([]entities.Product, error) {
	var res []entities.Product
	query := `SELECT * FROM PRODUCTS`
	if err := product.DB.Raw(query).Scan(&res).Error; err != nil {
		return nil, fmt.Errorf("error retrieving products")
	}
	return res, nil
}
func (product *ProductAdapter) GetProduct(id uint) (entities.Product, error) {
	var res entities.Product
	query := `SELECT * FROM products where id=$1`
	if err := product.DB.Raw(query, id).Scan(&res).Error; err != nil {
		return entities.Product{}, err
	}
	return res, nil
}
func (product ProductAdapter) IncrementQuantity(id uint, quantity int) (entities.Product, error) {
	var res entities.Product
	query := `UPDATE products SET quantity=$1 WHERE id=$2`
	if err := product.DB.Raw(query, quantity, id).Scan(&res).Error; err != nil {
		return entities.Product{}, err
	}
	return res, nil
}
func (product ProductAdapter) DecrementQuantity(id uint, quantity int) (entities.Product, error) {
	var res entities.Product
	query := `UPDATE products SET quantity=quantity-$1 WHERE id=$2`
	tx := product.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Raw(query, quantity, id).Scan(&res).Error; err != nil {
		tx.Rollback()
		return res, err
	}
	if res.Quantity < 0 {
		tx.Rollback()
		return res, fmt.Errorf("updated quantity can't be negative")
	}
	if err := tx.Commit().Error; err != nil {
		return res, err
	}
	return res, nil
}
