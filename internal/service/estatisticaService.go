package service

import (
	"desafioGolang/internal/model"
	"math"
	"time"
)

type EstatisticaService struct {
	transacaoService *TransacaoService
	janelaService    time.Duration
}

func NewEstatisticaService(transacaoService *TransacaoService, segundosConfig int) *EstatisticaService {
	return &EstatisticaService{
		transacaoService: transacaoService,
		janelaService:    time.Duration(segundosConfig) * time.Second,
	}
}

func (s *EstatisticaService) Calcular() model.EstatisticaResponse {
	todasTransacoes := s.transacaoService.Listar()
	var count int64
	var sum float64
	min := math.MaxFloat64
	max := -math.MaxFloat64

	agora := time.Now()

	for _, t := range todasTransacoes {
		
		tempoDecorrido := agora.Sub(t.DataHora)
		if tempoDecorrido >= 0 && tempoDecorrido <= s.janelaService {
			count++
			sum += t.Valor

			if t.Valor < min {
				min = t.Valor
			}
			if t.Valor > max {
				max = t.Valor
			}
		}
	}

	if count == 0 {
		return model.EstatisticaResponse{
			Count: 0,
			Sum:   0,
			Avg:   0,
			Min:   0.0,
			Max:   0.0,
		}
	}

	avg := sum / float64(count)

	return model.EstatisticaResponse{
		Count: count,
		Sum:   sum,
		Avg:   avg,
		Min:   min,
		Max:   max,
	}
}
