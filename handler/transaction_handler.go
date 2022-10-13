package handler

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/service"
	"assignment-golang-backend/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return

		}
		td.SourceID = int(uid)
		td.TargetID = int(uid)

		err = t.service.TopUpWallet(td)

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

func (t *TransactionHandler) TransferWallet() gin.HandlerFunc {
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

		targetId, err := utils.ExtractTokenID(ctx)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		td.SourceID = int(targetId)

		err = t.service.TransferWallet(td)

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

func (t *TransactionHandler) GetAllTransactionByLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		descReq := ctx.Request.URL.Query().Get("s")
		sortItem := ctx.Request.URL.Query().Get("sortBy")
		sortOrder := ctx.Request.URL.Query().Get("sort")
		limitNum := ctx.Request.URL.Query().Get("limit")

		descAny := fmt.Sprintf("%%%s%%", descReq)

		num, _ := strconv.Atoi(limitNum)

		NewTransactionRequest := entity.TransactionRequest{
			DescriptionRequest: descAny,
			SortByEntity:       sortItem,
			SortOrder:          sortOrder,
			Limit:              num,
		}

		fmt.Printf("NewTransactionRequest: %v\n", NewTransactionRequest)

		uid, err := utils.ExtractTokenID(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		tr, err := t.service.GetAllTransactionByLogin(int(uid), NewTransactionRequest)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"Data":    tr,
		})
	}

}
