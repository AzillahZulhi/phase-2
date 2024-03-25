package main

import (
	"fmt"
	"net/http"
	"preview-week-2/config"
	"preview-week-2/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.DbConnect("root:@tcp(localhost:3306)/preview_week_2")
	defer config.DB.Close()

	router := httprouter.New()

	router.GET("/branches", handler.GetAll)
	router.GET("/branches/:id", handler.GetByID)
	router.POST("/branches", handler.Create)
	router.PUT("/branches/:id", handler.Update)
	router.DELETE("/branches/:id", handler.Delete)

	fmt.Println("Running server on port :8083")
	err := http.ListenAndServe(":8083", router)
	if err != nil {
		fmt.Println("Error while starting server :", err.Error())
	}

}
