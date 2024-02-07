package service

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"github.com/vishnusunil243/ProductService/adapters"
	"github.com/vishnusunil243/ProductService/entities"
	"github.com/vishnusunil243/proto-files/pb"
)

var (
	Tracer opentracing.Tracer
)

func RetrieveTracer(tr opentracing.Tracer) {
	Tracer = tr
}

type ProductService struct {
	Adapter adapters.AdapterInterface
	pb.UnimplementedProductServiceServer
}

func NewProductService(adapter adapters.AdapterInterface) *ProductService {
	return &ProductService{
		Adapter: adapter,
	}
}
func (product *ProductService) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	span := Tracer.StartSpan("add products grpc")
	defer span.Finish()
	if req.Name == "" {
		return nil, fmt.Errorf("the name of the product can't be empty")
	}
	reqEntity := entities.Product{
		Name:     req.Name,
		Price:    uint(req.Price),
		Quantity: uint(req.Quantity),
	}
	res, err := product.Adapter.AddProduct(reqEntity)
	if err != nil {
		return nil, err
	}
	return &pb.AddProductResponse{
		Id:       uint32(res.Id),
		Name:     res.Name,
		Price:    int32(res.Price),
		Quantity: int32(res.Quantity),
	}, nil
}
func (product *ProductService) GetAllProducts(em *empty.Empty, srv pb.ProductService_GetAllProductsServer) error {
	span := Tracer.StartSpan("getall products grpc")
	defer span.Finish()
	products, err := product.Adapter.GetAllProducts()
	if err != nil {
		return err
	}
	for _, prod := range products {
		if err = srv.Send(&pb.AddProductResponse{
			Id:       uint32(prod.Id),
			Name:     prod.Name,
			Price:    int32(prod.Price),
			Quantity: int32(prod.Quantity),
		}); err != nil {
			return err
		}
	}
	return nil
}
