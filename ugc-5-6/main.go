package main

import (
	"fmt"
	"net/http"
	"ugc-5/config"
	"ugc-5/handler"
	"ugc-5/middleware"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.DbConnect("root:@tcp(localhost:3306)/avenger_corp_database")
	defer config.DB.Close()

	router := httprouter.New()

	router.POST("/register", handler.RegisterHandler)
	router.POST("/login", handler.LoginHandler)

	router.GET("/recipes/:id", middleware.AuthMiddleware(handler.GetByID))
	router.GET("/recipes", middleware.AuthMiddleware(handler.GetAll))
	router.PUT("/recipes/:id", middleware.AuthMiddleware(handler.Update))

	//superAdmin
	router.POST("/recipes", middleware.AuthMiddleware(middleware.AdminOnlyMiddleware(handler.Create)))
	router.DELETE("/recipes/:id", middleware.AuthMiddleware(middleware.AdminOnlyMiddleware(handler.Delete)))

	fmt.Println("Running server on port :8082")

	err := http.ListenAndServe(":8082", router)
	if err != nil {
		fmt.Println("Error while starting server :", err.Error())
	}
}
