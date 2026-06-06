package service

import (
	"desafioGolang/internal/model"
	"errors"
	"sync"
	"time"
)

var (
	ErrValorNegativo = errors.New("o valor da transação não pode ser negativo")
	ErrDataFutura    = errors.New("a data da transacao nao pode ser no futuro")
)

type TransacaoService struct {
	mu        sync.RWMutex
	trasacoes []model.TransacaoRequest
}

func NewTransacaoService() *TransacaoService {
	return &TransacaoService{
		trasacoes: make([]model.TransacaoRequest, 0),
	}
}

func (t *TransacaoService) AddTrasaco(req model.TransacaoRequest) error {
	if req.Valor < 0 {
		return ErrValorNegativo
	}
	if req.DataHora.After(time.Now()) {
		return ErrDataFutura
	}
	t.mu.Lock()
	defer t.mu.Unlock()

	t.trasacoes = append(t.trasacoes, req)
	return nil
}

func (t *TransacaoService) Limpar() {
	t.mu.Lock()
	t.mu.Unlock()
	t.trasacoes = make([]model.TransacaoRequest, 0)
}

func (t *TransacaoService) Listar() []model.TransacaoRequest {
	t.mu.RLock()
	defer t.mu.RUnlock()

	copia := make([]model.TransacaoRequest, len(t.trasacoes))
	copy(copia, t.trasacoes)
	return copia
}
