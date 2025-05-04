package handlers

import (
	"net/http"

	"github.com/Cayo-Cezar/go-fruits-api/services"
	"github.com/gin-gonic/gin"
)

type FruitHandler struct {
	Service *services.FruitService
}

func NewFruitHandler(service *services.FruitService) *FruitHandler {
	return &FruitHandler{Service: service}
}

func (h *FruitHandler) GetSugarReport(c *gin.Context) {
	report, err := h.Service.GenerateSugarReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar relatório"})
		return
	}
	c.JSON(http.StatusOK, report)
}
