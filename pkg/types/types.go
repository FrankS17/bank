package types


// Представляет собой денежную сумму в минимальных единицах (копейки и т.д.)
type Money int64


// Category представляет собой категорию, в которой был совершен платеж
type Category string

type Ron string

type Status string

const (
	StatusOk Status = "Ok"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct {
	ID int 
	Amount Money
	Category Category
	Status Status
	Q2 Ron
}


