package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}
func (b *bill) format() string {
	fs := "Bill BreakDown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v...RS%v \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-25v...Rs%v\n", "Tip: ", b.tip)
	fs += fmt.Sprintf("%-25v...Rs%0.2f", "total:", total+b.tip)
	return fs
}
func (b *bill) update(tip float64) {
	b.tip = tip
}
func (b *bill) updateitems(name string, price float64) {
	b.items[name] = price
}

func (b bill) save() {
	data := []byte(b.format())

	err := os.WriteFile(b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
