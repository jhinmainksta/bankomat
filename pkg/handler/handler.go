package handler

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jhinmainksta/bankomat/pkg/service"
)

type Handler struct {
	accounts []service.BankAccount
	mu       sync.Mutex
}

func NewHandler(accounts []service.BankAccount) *Handler {
	return &Handler{accounts: accounts}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	acc := router.Group("/accounts")
	{
		acc.POST("/", h.create)
		acc.POST("/:id/deposit", h.deposit)
		acc.POST("/:id/withdraw", h.withdraw)
		acc.GET("/:id/balance", h.balance)
	}

	return router
}
