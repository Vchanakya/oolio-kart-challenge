package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vchanakya/oolio-kart-challenge/backend/api"
	"github.com/Vchanakya/oolio-kart-challenge/backend/db"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/handler"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/repository"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/seed"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/utils"
)

func main() {
	server := &api.Server{}
	database, err := db.ConnectDB("oolio")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")
	seed.Seed("oolio")
	fmt.Println("Database seeded")
	handler := handler.NewHandler(server, repository.NewRepository(database))
	srv, err := api.NewServer(handler, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", utils.KeyMiddleware(srv)); err != nil {
		log.Fatal(err)
	}
}
