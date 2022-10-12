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

type TransactionHandler struct {
	service service.TransactionServices
}

func NewTransactionHandler(service service.TransactionServices) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (t *TransactionHandler) TopUpWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		var td entity.Transaction
		err = json.Unmarshal(reqBody, &td)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		uid, err := utils.ExtractTokenID(ctx)
		td.SourceID = int(uid)
		td.TargetID = int(uid)
		//fmt.Printf("ul di handler: %v\n", ul)
		t.service.TopUpWallet(td)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"Data":    td,
		})

	}
}
