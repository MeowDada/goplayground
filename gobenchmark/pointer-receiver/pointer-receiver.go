package benchmark

type normalStruct struct {
	a int
	b int
	c int
}

func (n normalStruct) Add() int {
	return 1
}

type ptrRecvStruct struct {
	a int
	b int
	c int
}

func (p *ptrRecvStruct) Add() int {
	return 1
}