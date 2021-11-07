package types


// Представляет собой денежную сумму в минимальных единицах (копейки и т.д.)
type Money int64


// Category представляет собой категорию, в которой был совершен платеж
type Category string

type Ron string

type Payment struct {
	ID int 
	Amount Money
	Category Category
	Q2 Ron
}


