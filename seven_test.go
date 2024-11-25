package vv26s

import (
	"errors"
	"testing"
)

var (
	// ValidTestCasesLessEqualThousand possui os casos de teste com produtos válidos cuja fatura sem desconto possui valor menor ou igual a 1000
	ValidTestCasesLessEqualThousand = []TestCase{
		// caso de teste com valor total menor que 1000
		{
			Products: []Product{
				{
					Name:     "Product 1",
					Price:    150,
					Quantity: 3,
				},
				{
					Name:     "Product 2",
					Price:    30,
					Quantity: 2,
				},
			},
			Discount: 25,
			Expected: 382.5,
		},
		// caso de teste com valor total igual a 1000
		{
			Products: []Product{
				{
					Name:     "Product 1",
					Price:    150,
					Quantity: 6,
				},
				{
					Name:     "Product 4",
					Price:    100,
					Quantity: 1,
				},
			},
			Discount: 15,
			Expected: 850,
		},
		// caso de teste com valor total menor que 1000 e desconto 0
		{
			Products: []Product{
				{
					Name:     "Product 5",
					Price:    75,
					Quantity: 2,
				},
				{
					Name:     "Product 6",
					Price:    15,
					Quantity: 1,
				},
			},
			Discount: 0,
			Expected: 165,
		},
	}

	// ValidTestCasesGreaterThanThousand possui os casos de teste com produtos válidos cuja fatura sem desconto possui valor maior que 1000
	ValidTestCasesGreaterThanThousand = []TestCase{
		// caso de teste com valor total maior que 1000, com valores quebrados
		{
			Products: []Product{
				{
					Name:     "Product 1",
					Price:    150,
					Quantity: 5,
				},
				{
					Name:     "Product 3",
					Price:    67.25,
					Quantity: 4,
				},
			},
			Discount: 17.75,
			Expected: 755.88,
		},
		// caso de teste com valor total maior que 1000 e desconto 0
		{
			Products: []Product{
				{
					Name:     "Product 5",
					Price:    75,
					Quantity: 20,
				},
				{
					Name:     "Product 6",
					Price:    15,
					Quantity: 1,
				},
			},
			Discount: 0,
			Expected: 1415,
		},
		// caso de teste com valor total maior que 1000 e desconto negativo
		// desconto negativo não foi considerado na prova, então teoricamente é 'válido'
		// vai acrescer no valor
		// e como é maior que 1000, vai primeiro descontar 100 antes de aplicar o 'desconto'
		{
			Products: []Product{
				{
					Name:     "Product 5",
					Price:    75,
					Quantity: 20,
				},
				{
					Name:     "Product 6",
					Price:    15,
					Quantity: 1,
				},
			},
			Discount: -10,
			Expected: 1556.5,
		},
	}

	// InvalidTestCases possui os casos de teste com produtos inválidos
	// ou seja, produtos com preço ou quantidade com valor negativo
	InvalidTestCases = []TestCase{
		// caso de teste com preço negativo
		{
			Products: []Product{
				{
					Name:     "Product -1",
					Price:    -150,
					Quantity: 3,
				},
				{
					Name:     "Product 2",
					Price:    30,
					Quantity: 2,
				},
			},
			Discount: 25,
			Expected: 0, // vai falhar, então não tem valor esperado
		},
		// caso de teste com quantidade negativa
		{
			Products: []Product{
				{
					Name:     "Product 1",
					Price:    150,
					Quantity: 3,
				},
				{
					Name:     "Product 2",
					Price:    30,
					Quantity: -2,
				},
			},
			Discount: 10,
			Expected: 0, // vai falhar, então não tem valor esperado
		},
	}
)

// TestInvoiceDiscountOnValidProducts testa os casos em que os produtos são válidos
// ou seja, os produtos possuem preço e quantidade com valor positivo
func TestInvoiceDiscountOnValidProducts(t *testing.T) {
	t.Log("testing valid products with gross total invoice value less or equal to 1000")
	for _, validTestCase := range ValidTestCasesLessEqualThousand {
		totalInvoiceValue, err := CalculateInvoice(validTestCase.Products, validTestCase.Discount)
		if err != nil {
			t.Errorf("valid test case <=1000 failed; got error: %v", err)
		}

		if totalInvoiceValue != validTestCase.Expected {
			t.Errorf("valid test case <=1000 failed; expected %v, got %v", validTestCase.Expected, totalInvoiceValue)
		}
	}

	t.Log("testing valid products with gross total invoice value greater than 1000")
	for _, validTestCase := range ValidTestCasesGreaterThanThousand {
		totalInvoiceValue, err := CalculateInvoice(validTestCase.Products, validTestCase.Discount)
		if err != nil {
			t.Errorf("valid test case >1000 failed; got error: %v", err)
		}

		if totalInvoiceValue != validTestCase.Expected {
			t.Errorf("valid test case >1000 failed; expected %v, got %v", validTestCase.Expected, totalInvoiceValue)
		}
	}
}

// TestInvoiceDiscountOnInvalidProducts testa os casos em que os produtos são inválidos
// ou seja, produtos com preço ou quantidade com valor negativo
func TestInvoiceDiscountOnInvalidProducts(t *testing.T) {
	t.Log("testing invalid products")
	for _, invalidTestCase := range InvalidTestCases {
		totalInvoiceValue, err := CalculateInvoice(invalidTestCase.Products, invalidTestCase.Discount)
		if err == nil {
			t.Errorf("invalid test case failed; expected error, got %v", totalInvoiceValue)
		}
		if !errors.Is(err, ErrInvalidProduct) {
			t.Errorf("invalid test case failed; expected ErrInvalidProduct, got %v", err)
		}
	}
}
