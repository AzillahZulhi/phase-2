package main

import (
	"fmt"
	"net/http"
	"ugc-2/config"
	"ugc-2/handler"
)

func main() {
	db := config.ConnectDB("root:@tcp(localhost:3306)/superhero_database")
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/get/heroes", func(w http.ResponseWriter, r *http.Request) {
		handler.GetHeroesData(w, r)
	})
	mux.HandleFunc("/get/villain", func(w http.ResponseWriter, r *http.Request) {
		handler.GetVillainsData(w, r)
	})

	fmt.Println("Running server on port :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error while starting server :", err.Error())
	}
}
