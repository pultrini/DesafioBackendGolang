package service

import (
	"desafioGolang/internal/model"
	"sync"
	"testing"
	"time"
)

func TestCalcularEstatistica(t *testing.T) {
	ts := NewTransacaoService()
	es := NewEstatisticaService(ts, 60)

	t.Run("Calulo apenas dos 60s anteriores", func(t *testing.T) {
		ts.AddTrasaco(model.TransacaoRequest{
			Valor:    20.5,
			DataHora: time.Now(),
		})
		ts.AddTrasaco(model.TransacaoRequest{
			Valor:    5.5,
			DataHora: time.Now(),
		})

		ts.AddTrasaco(model.TransacaoRequest{
			Valor:    1000,
			DataHora: time.Now().Add(-120 * time.Second),
		})

		stats := es.Calcular()

		if stats.Count != 2 {
			t.Errorf("Era esperado dois mas houve %d", stats.Count)
		}
		if stats.Avg != 13 {
			t.Errorf("Era esperado ter 13 de média mas houve %f", stats.Avg)
		}
		if stats.Max != 20.5 {
			t.Errorf("era esperado ter 20.5 de mas houve %f", stats.Max)
		}
		if stats.Min != 5.5 {
			t.Errorf("era esperado ter 5.5 de mas houve %f", stats.Min)
		}
	})
}

func TestSegurancaDeConcorrencia(t *testing.T) {
	ts := NewTransacaoService()
	es := NewEstatisticaService(ts, 60)

	var wg sync.WaitGroup
	totalRequisicoes := 100

	wg.Add(totalRequisicoes)
	for i := 0; i < totalRequisicoes; i++ {
		go func(valor float64) {
			defer wg.Done()
			ts.AddTrasaco(model.TransacaoRequest{Valor: valor, DataHora: time.Now()})
			es.Calcular()
		}(float64(i))
	}
	wg.Wait()
	stats := es.Calcular()
	if stats.Count != int64(totalRequisicoes) {
		t.Errorf("Falha de concorrencia: Esperava %d, registrou %d", totalRequisicoes, stats.Count)
	}
}
