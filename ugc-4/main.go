package main

import (
	"fmt"
	"net/http"
	"ugc-4/config"
	"ugc-4/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.DbConnect("root:@tcp(localhost:3306)/avenger_corp_database")
	defer config.DB.Close()

	router := httprouter.New()

	router.GET("/criminal_report", handler.GetALLCriminalReport)
	router.GET("/criminal_report/:id", handler.GetCriminalReport)
	router.POST("/criminal_report", handler.CreateCriminalReport)
	router.PUT("/criminal_report/:id", handler.UpdateCriminalReport)
	router.DELETE("/criminal_report/:id", handler.DeleteReport)

	fmt.Println("Running server on port :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error while starting server :", err.Error())
	}
}
