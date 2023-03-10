package main

import (
	"assignment-golang-backend/handler"
	"assignment-golang-backend/pkg/database/postgres"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/service"
	"assignment-golang-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := postgres.DbInit()
	router := gin.Default()

	router.Static("/docs", "./dist")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserServices(userRepository)
	userHandler := handler.NewUserHandler(userService)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionServices(transactionRepository, userRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Url not valid"})
	})

	router.POST("/register", userHandler.RegisterUser())
	router.POST("/login", userHandler.LoginUser())
	router.Use(utils.JwtAuthMiddleware()).GET("/user-detail", userHandler.GetUserDetails())

	router.Use(utils.JwtAuthMiddleware()).POST("/top-up", transactionHandler.TopUpWallet())
	router.Use(utils.JwtAuthMiddleware()).POST("/transfer", transactionHandler.TransferWallet())
	router.Use(utils.JwtAuthMiddleware()).GET("/transactions", transactionHandler.GetAllTransactionByLogin())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "assignment-go-be running",
		})
	})
	router.Run()
}
