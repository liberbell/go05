package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	// t1 := Trade{"MSFT", 10, 99.98, true}
	// fmt.Println(t1)
	//
	// fmt.Printf("%+v\n", t1)
	// fmt.Println(t1.Symbol)
	//
	// t2 := Trade{
	// 	Symbol: "MSFT",
	// 	Volume: 10,
	// 	Price:  99.98,
	// 	Buy:    true,
	// }
	// fmt.Printf("%+v\n", t2)
	//
	// t3 := Trade{}
	// fmt.Printf("%+v\n", t3)

	t := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Println(t.Value())
}
