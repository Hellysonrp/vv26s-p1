package vv26s

import "errors"

var (
	// ErrInvalidProduct é equivalente ao InvalidProductException requisitado na prova
	//
	// golang não possui exceções, somente erros
	//
	// foi a segunda coisa a ser escrita, para referenciar nos testes
	ErrInvalidProduct = errors.New("invalid product")
)

// Product é o tipo que representa um produto nos testes
//
// foi a primeira coisa a ser escrita, para conseguir referenciar nos testes
type Product struct {
	Name     string
	Price    float64
	Quantity int64
}

// TestCase representa um caso de teste
//
// quarta coisa a ser escrita...
type TestCase struct {
	Products []Product
	Discount float64
	Expected float64
}
