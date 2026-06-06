package handler

import (
	"desafioGolang/internal/model"
	"desafioGolang/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransacaoHandler struct {
	transacaoService *service.TransacaoService
}

func NewTransacaoHandler(ts *service.TransacaoService) *TransacaoHandler {
	return &TransacaoHandler{
		transacaoService: ts,
	}
}

func (h *TransacaoHandler) PostTransacao(c *gin.Context) {
	var req model.TransacaoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON mal formatado ou invalido"})
		return
	}
	err := h.transacaoService.AddTrasaco(req)
	if err != nil {
		if errors.Is(err, service.ErrValorNegativo) || errors.Is(err, service.ErrDataFutura) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
	}
	c.Status(http.StatusCreated)
}

func (h *TransacaoHandler) Limpar(c *gin.Context) {
	h.transacaoService.Limpar()
	c.Status(http.StatusOK)
}
