package main

import ( 

	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/ncostamagna/go-sp-products/bootstrap"
	"github.com/ncostamagna/go-sp-products/adapter/postgres"
	"github.com/ncostamagna/go-sp-products/internal/product"
	"github.com/ncostamagna/go-sp-products/transport/httpapi"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := bootstrap.InitPostgres()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := postgres.NewRepository(db)
	
	srv := product.NewService(repo)

	endpoints := httpapi.MakeProductsEndpoints(srv)
	apiServer := httpapi.New(endpoints)

	errs := make(chan error, 1)

	go func() {
		url := fmt.Sprintf("0.0.0.0:8086")
		log.Println("Listening", "url", url)
		errs <- apiServer.Run(url)
	}()

	fatalErr := <-errs
	log.Println("Program ended:", fatalErr)
}