package handler

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/service"
	"assignment-golang-backend/utils"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserServices
}

func NewUserHandler(service service.UserServices) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		var ul entity.User
		err = json.Unmarshal(reqBody, &ul)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = u.service.RegisterUser(ul)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"Data":    ul,
		})

	}
}

func (u *UserHandler) LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		var ul entity.User
		err = json.Unmarshal(reqBody, ul)

		err = u.service.LoginUser(ul.Email, ul.Password)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
		}

		token, _ := utils.GenerateToken(uint(ul.ID))
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"Token":   token,
		})

	}
}

func (u *UserHandler) GetAllUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// params := ctx.Request.URL.Query().Get("name")

		// queryUser := fmt.Sprintf("%%%s%%", params)
		// fmt.Printf("queryUser: %v\n", queryUser)

		ul, err := u.service.GetAllUser()
		if err != nil {

			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"Data":    ul,
		})
	}

}
