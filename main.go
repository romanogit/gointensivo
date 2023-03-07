package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Pessoa struct {
	nome string
}

func NewPessoa() *Pessoa {
	return &Pessoa{}
}

func main() {
	fmt.Println(uuid.New().String())
	
	a := NewPessoa()
	a.nome = "Romano"
	b := a
	c := NewPessoa()
	c.nome = "Silva"

	fmt.Println(a.nome)
	fmt.Println(b.nome)
	fmt.Println(c.nome)
}