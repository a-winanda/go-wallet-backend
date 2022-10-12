package main

import (
	"assignment-golang-backend/handler"
	"assignment-golang-backend/pkg/database/postgres"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := postgres.DbInit()
	router := gin.Default()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserServices(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.POST("/register", userHandler.RegisterUser())
	router.POST("/login", userHandler.LoginUser())
	router.GET("/users", userHandler.GetAllUser())

	//router.Static("/documentation", "dist/")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "exercise-gin running",
		})
	})
	router.Run()
}
