package model

import (
	"time"
)

type TransacaoRequest struct {
	Valor    float64   `json:"valor"`
	DataHora time.Time `json:"dataHora"`
}
