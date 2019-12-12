package main

import (
	"fmt"
)

type myStruct struct {
	id int
	intArr []int
	squareMap map[int]int
}

func (m myStruct) Modify1() {
	fmt.Printf("%-26s = %p\n", "mem addr of ms", &m)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.id", &m.id)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.intArr", m.intArr)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.squareMap", m.squareMap)
	m.id = m.id + 1
	for idx, _ := range m.intArr {
		m.intArr[idx] = m.intArr[idx]+1
	}
	for k, _ := range m.squareMap {
		m.squareMap[k] = (k+2)*(k+2)
	}
}

func (m *myStruct) Modify2() {
	fmt.Printf("%-26s = %p\n", "mem addr of ms", m)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.id", &m.id)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.intArr", m.intArr)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.squareMap", m.squareMap)
	m.id = m.id + 1
	for idx, _ := range m.intArr {
		m.intArr[idx] = m.intArr[idx]+1
	}
	for k, _ := range m.squareMap {
		m.squareMap[k] = (k+3)*(k+3)
	}
}

func newMyStruct(id int, args ...int) myStruct {
	intArr := args
	squareMap := make(map[int]int)
	for _, num := range intArr {
		squareMap[num] = num*num
	}
	return myStruct{
		id: id,
		intArr: intArr,
		squareMap: squareMap,
	}
}

func tryModifyMyStructInternalDataByCopy(ms myStruct) {
	fmt.Printf("%-26s = %p\n", "mem addr of ms", &ms)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.id", &ms.id)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.intArr", ms.intArr)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.squareMap", ms.squareMap)
	
	ms.id = ms.id + 1

	for idx, _ := range ms.intArr {
		ms.intArr[idx] = ms.intArr[idx] + 1 
	}

	for k, _ := range ms.squareMap {
		ms.squareMap[k] = (k+1)*(k+1)
	}
}

func main() {

	// Creat a struct and dump its info.
	ms := newMyStruct(0,1,2,3,4,5)
	fmt.Printf("%-26s = %p\n", "mem addr of ms", &ms)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.id", &ms.id)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.intArr", ms.intArr)
	fmt.Printf("%-26s = %p\n", "mem addr of ms.squareMap", ms.squareMap)
	fmt.Println(ms)

	// Pass the struct as value into a function.
	// And this function will try modify internal data.
	tryModifyMyStructInternalDataByCopy(ms)
	fmt.Println(ms)

	// Result1: All value except id has been changed succefully.
	// Although the struct arguemnt we pass into the function will be a copied one.
	// But the array and the map inside the copied one points to the same memory address
	// as the original one we pass into it. So the array and the map could be modified.
	// However, the id field is not a pointer, it just a copy. so modify this value won't
	// do any affect to the original data structure.

	// Calling struct method and check the internal data will change or not.
	ms.Modify1()
	fmt.Println(ms)

	// Result2: Exactly same as Result1.

	// Calling pointer receiver method and check the internal data will change or not.
	ms.Modify2()
	fmt.Println(ms)

	// Result3: All value is the original one. so everything could be modified as we expect.
}