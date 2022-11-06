package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create new Bill name : ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create new Bill name : ", reader)
	b := newBill(name)
	fmt.Println("created the bill -", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Chose Option(a-add item ,s-save bill ,t-add tips) : ", reader)

	switch opt {
	case "a":
		name, _ := getInput("what is the item name: ", reader)
		price, _ := getInput("what is the item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price should be number")
			promptOptions(b)
		}
		b.updateitems(name, p)
		fmt.Println("Items added", name, price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("U chose to saved the bill", b.name)

	case "t":
		tip, _ := getInput("what is the tip Ypu will give: ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip should be number")
			promptOptions(b)
		}
		b.update(t)
		fmt.Println("tip added", t)
		promptOptions(b)

	default:
		fmt.Println("That was not valid")
		promptOptions(b)
	}

}
func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}
