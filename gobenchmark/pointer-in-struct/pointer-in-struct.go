package benchmark

type ptrRecvStruct struct {
	a int
	b int
	c int
}

func (s *ptrRecvStruct) changeValue(a, b, c int) {
	s.a, s.b, s.c = a, b, c
}

type ptrInStruct struct {
	a *int
	b *int
	c *int
}

func (s *ptrInStruct) changeValue(a, b, c *int) {
	s.a, s.b, s.c = a, b, c
}