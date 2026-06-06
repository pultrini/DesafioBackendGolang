package handler

import (
	"desafioGolang/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EstatisticaHandler struct {
	estatisticaService *service.EstatisticaService
}

func NewEstatisticaHandel(es *service.EstatisticaService) *EstatisticaHandler {
	return &EstatisticaHandler{estatisticaService: es}
}

func (h *EstatisticaHandler) GetEstatistica(c *gin.Context) {
	res := h.estatisticaService.Calcular()
	c.JSON(http.StatusOK, res)
}
