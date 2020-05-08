package composite

type CopyStruct struct {
	val int
	str string
}

func (c CopyStruct) SetVal(val int) CopyStruct {
	c.val = val
	return c
}

func (c CopyStruct) SetString(str string) CopyStruct {
	c.str = str
	return c
}

type PtrStruct struct {
	val int
	str string
}

func (p *PtrStruct) SetVal(val int) *PtrStruct {
	p.val = val
	return p
}

func (p *PtrStruct) SetString(str string) *PtrStruct {
	p.str = str
	return p
}