package main

import (
	"encoding/json"
	myjsonobj "github.com/MeowDada/goplayground/gojson/pkg/jsonobj"
	"fmt"
)

func main() {

	toyota := myjsonobj.Car {
		Name:   "RX001-5",
		Brand:  "Toyota",
		Price:  myjsonobj.PriceInfo {
			Value:        30000,
			Concurrency: "USD",
		},
		Owner: myjsonobj.Person {
			Name:   "Jack",
			Age:    26,
			Gender: "Male",
		},
	}

	bytes, err := json.MarshalIndent(toyota, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Marshal result:")
	fmt.Println(string(bytes))

	var car myjsonobj.Car
	if err := json.Unmarshal(bytes, &car); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Unmarshal result:")
	fmt.Println(car)
}