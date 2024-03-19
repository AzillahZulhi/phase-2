package main

import (
	"fmt"
	"net/http"
	"ugc-3/config"
	"ugc-3/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.DbConnect("root:@tcp(localhost:3306)/avenger_corp_database")
	defer config.DB.Close()

	router := httprouter.New()

	router.GET("/inventories", handler.GetAll)
	router.GET("/inventories/:id", handler.GetInventory)
	router.POST("/inventories", handler.Create)
	router.PUT("/inventories/:id", handler.Update)
	router.DELETE("/inventories/:id", handler.Delete)

	fmt.Println("Running server on port :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error while starting server :", err.Error())
	}
}
