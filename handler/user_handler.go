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
			return
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
		err = json.Unmarshal(reqBody, &ul)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		//fmt.Printf("ul di handler: %v\n", ul)

		token, err := u.service.LoginUser(ul.Email, ul.Password)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"Token":   token,
		})

	}
}

func (u *UserHandler) GetUserDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		uid, err := utils.ExtractTokenID(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		wallet, user, err := u.service.GetUserDetails(int(uid))
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"Data": gin.H{
				"user_id": user.ID,
				"email":   user.Email,
				"wallet": gin.H{
					"number":  wallet.WalletNumber,
					"balance": wallet.Balance,
				},
			},
		})

	}
}
