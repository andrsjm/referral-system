package main

import (
	"fmt"
	"log"
	"net/http"
	"referral-system/app"

	"github.com/rs/cors"
)

func main() {
	router := app.SetupRouter()

	corsHandler := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(router)

	http.Handle("/", handler)
	fmt.Println("Connected to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
