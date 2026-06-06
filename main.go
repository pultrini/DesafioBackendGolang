package main

import (
	"desafioGolang/internal/handler"
	"desafioGolang/internal/service"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	segundosConfig := 60
	if envVar := os.Getenv("SEGUNDOS_CONFIG"); envVar != "" {
		if val, err := strconv.Atoi(envVar); err == nil {
			segundosConfig = val
		}
	}
	transacaoService := service.NewTransacaoService()
	estatisticaService := service.NewEstatisticaService(transacaoService, segundosConfig)

	transacaoHandler := handler.NewTransacaoHandler(transacaoService)
	estatisticaHandler := handler.NewEstatisticaHandel(estatisticaService)

	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/transacao", transacaoHandler.PostTransacao)
	r.DELETE("/transcao", transacaoHandler.Limpar)
	r.GET("/estatistica", estatisticaHandler.GetEstatistica)

	println("[Servidor] API rodando na porta 8080")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
