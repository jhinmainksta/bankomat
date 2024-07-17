package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jhinmainksta/bankomat/pkg/service"
)

func (h *Handler) create(c *gin.Context) {
	account := service.Account{}
	h.accounts = append(h.accounts, &account)

	succesResponse(c, http.StatusCreated, fmt.Sprintf("%d bank account has been created", len(h.accounts)-1), map[string]interface{}{
		"id": len(h.accounts) - 1,
	})

}

func (h *Handler) deposit(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = c.Request.ParseForm()
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	am := service.Amount{0}
	if err := c.BindJSON(&am); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.mu.Lock()
	err = h.accounts[id].Deposit(am.Amount)
	h.mu.Unlock()

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	succesResponse(c, http.StatusCreated, fmt.Sprintf("succesfuly deposed %f рублей on %d bank account", am.Amount, id), statusResp{"ok"})

}

func (h *Handler) withdraw(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	am := service.Amount{0}
	if err := c.BindJSON(&am); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.mu.Lock()
	err = h.accounts[id].Withdraw(am.Amount)
	h.mu.Unlock()

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	succesResponse(c, http.StatusCreated, fmt.Sprintf("succesfuly withdrawed %f рублей on %d bank account", am.Amount, id), statusResp{"ok"})

}

func (h *Handler) balance(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	ans := service.Amount{h.accounts[id].GetBalance()}

	succesResponse(c, http.StatusOK, fmt.Sprintf("balance on %d bank account is %f", id, ans.Amount), ans)

}
