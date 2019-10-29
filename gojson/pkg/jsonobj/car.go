package jsonobj

type PriceInfo struct {
	Value int
	Concurrency string
}

type Car struct {
	Name  string
	Brand string
	Price PriceInfo
	Owner Person
}