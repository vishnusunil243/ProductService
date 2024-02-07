package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	jaegar "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/vishnusunil243/ProductService/db"
	"github.com/vishnusunil243/ProductService/initializer"
	"github.com/vishnusunil243/ProductService/service"
	"github.com/vishnusunil243/proto-files/pb"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf(err.Error())
	}
	addr := os.Getenv("DB_KEY")

	DB, err := db.InitDB(addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	services := initializer.Initializer(DB)
	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, services)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen on port 8080: %v ", err)
	}
	log.Printf("product server listening on port 8080")
	tracer, closer := initTracer()

	defer closer.Close()
	service.RetrieveTracer(tracer)
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to listen on port 8080: %v", err)
	}
}
func initTracer() (tracer opentracing.Tracer, closer io.Closer) {
	jargarEndpoint := "http://localhost:14268/api/traces"
	cfg := &config.Configuration{
		ServiceName: "product-service",
		Sampler: &config.SamplerConfig{
			Type:  jaegar.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: jargarEndpoint,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("updated")
	return
}
