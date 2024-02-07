package initializer

import (
	"github.com/vishnusunil243/ProductService/adapters"
	"github.com/vishnusunil243/ProductService/service"
	"gorm.io/gorm"
)

func Initializer(db *gorm.DB) *service.ProductService {
	adapter := adapters.NewProductAdapter(db)
	service := service.NewProductService(adapter)
	return service
}
