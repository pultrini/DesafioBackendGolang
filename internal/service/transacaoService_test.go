package service

import (
	"desafioGolang/internal/model"
	"errors"
	"testing"
	"time"
)

func TestAdicionarTransacao(t *testing.T) {
	svc := NewTransacaoService()

	t.Run("Deve adicioonar transacao Valida", func(t *testing.T) {
		req := model.TransacaoRequest{
			Valor:    10.50,
			DataHora: time.Now(),
		}
		err := svc.AddTrasaco(req)
		if err != nil {
			t.Errorf("Erro nao esperado %v", err)
		}
	})

	t.Run("Não pode valor negativo", func(t *testing.T) {
		req := model.TransacaoRequest{
			Valor:    -10.50,
			DataHora: time.Now(),
		}
		err := svc.AddTrasaco(req)
		if !errors.Is(err, ErrValorNegativo) {
			t.Errorf("Erro diferente do formatado %v", err)

		}
	})

	t.Run("Deve Barrar data futura", func(t *testing.T) {
		req := model.TransacaoRequest{
			Valor:    10.50,
			DataHora: time.Now().AddDate(0, 0, 1),
		}
		err := svc.AddTrasaco(req)
		if !errors.Is(err, ErrDataFutura) {
			t.Errorf("Erro diferente do formatado %v", err)
		}
	})

}
func TestLimparTransacao(t *testing.T) {
	svc := NewTransacaoService()
	svc.AddTrasaco(model.TransacaoRequest{Valor: 10.50, DataHora: time.Now()})
	svc.Limpar()

	lista := svc.Listar()
	if len(lista) != 0 {
		t.Errorf("A lista deveria ser vazia e veio com %d itens", len(lista))
	}
}
