package vv26s

import (
	"math"
)

// CalculateInvoice calcula o total da fatura com o devido desconto aplicado
//
// products é a lista de produtos a ser considerada;
// discount é o percentual de desconto a ser aplicado (ex: 10 seria 10%)
//
// com base no valor total da fatura antes de aplicar o desconto, se o valor ultrapassar 1000, um desconto adicional de valor 100 será aplicado
//
// retorna o valor total da fatura após aplicar o desconto ou ErrInvalidProduct, caso algum produto possua preço ou quantidade negativa;
// o valor retornado é arredondado para duas casas decimais utilizando math.Round
//
// a assinatura desta função foi a terceira coisa a ser escrita, para poder referenciar nos testes
func CalculateInvoice(products []Product, discount float64) (totalInvoiceValue float64, err error) {
	for _, product := range products {
		if product.Price < 0 || product.Quantity < 0 {
			return 0, ErrInvalidProduct
		}

		totalInvoiceValue += product.Price * float64(product.Quantity)
	}

	// arredonda a primeira vez antes de checar se é maior que 1000 por conta de erros de precisão do float64
	totalInvoiceValue = math.Round(totalInvoiceValue*100) / 100

	if totalInvoiceValue > 1000 {
		totalInvoiceValue -= 100
	}

	totalInvoiceValue = math.Round(totalInvoiceValue*(1-discount/100)*100) / 100

	return
}
