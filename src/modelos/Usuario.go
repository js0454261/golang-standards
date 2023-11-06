package modelos

import "time"

//usuario representa um usuario utilizando nosso sistema
type usuario struct {
	Id       uint64    `json:"id"`
	Nome     string    `json:"nome"`
	Cpf      string    `json:"cpf,omitempty"` //quando temos esse omitempty é para declarar que não aceito esse campo em vazio
	Nick     string    `json:"nick"`
	Email    string    `json:"email"`
	Senha    string    `json:"senha"`
	CriadoEm time.Time `json:"criadoem"`
}
