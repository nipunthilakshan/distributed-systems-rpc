package main

import (
	"fmt"
	"net/rpc"

	"rpc-demo/common"
)

func main() {

	getPriceOfGivenVegetable()
	getAmountOfGivenVegetable()
	getAllVegetables()
	//addNewVegetable()
	//updateVegetable()

}

func getPriceOfGivenVegetable() {
	var vegName string
	fmt.Println("Enter vegetable Name to get the price per kg: ")
	// Taking input from user
	_, _ = fmt.Scanln(&vegName)

	// get RPC client by dialing at `rpc-demo.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	// create veg variable of type `common.Vegetable`
	var veg common.Vegetable
	if err := client.Call("Vegetables.Get", vegName, &veg); err != nil {
		fmt.Println("Error:1 Vegetables.Get()", err)
	} else {
		fmt.Printf("Success: Price per kg of '%s' is '%0.0f' \n", vegName, veg.Price)
	}

	fmt.Println("---------------------- End of Operation -------------------------")
}

func getAmountOfGivenVegetable() {
	var vegName string
	fmt.Println("Enter vegetable Name to get the available amount in kg: ")
	// Taking input from user
	_, _ = fmt.Scanln(&vegName)

	// get RPC client by dialing at `rpc-demo.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	// create veg variable of type `common.Vegetable`
	var veg common.Vegetable
	if err := client.Call("Vegetables.Get", vegName, &veg); err != nil {
		fmt.Println("Error:1 Vegetables.Get()", err)
	} else {
		fmt.Printf("Success: Amount of kg for '%s' is '%0.0f' \n", vegName, veg.Amount)
	}
	fmt.Println("---------------------- End of Operation -------------------------")
}

func getAllVegetables() {
	// get RPC client by dialing at `rpc-demo.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	// create veg variable of type `common.Vegetable`
	var vegNames []string
	if err := client.Call("Vegetables.GetAll", "dummy", &vegNames); err != nil {
		fmt.Println("Error:1 Vegetables.GetAll()", err)
	} else {
		fmt.Print("Success: Names of all vegetables: ")
		for i := 0; i < len(vegNames); i++ {
			fmt.Print(vegNames[i] + ", ")
		}
		fmt.Println("")
	}
	fmt.Println("---------------------- End of Operation -------------------------")
}

func addNewVegetable() {
	var vegName string
	var vegPrice float64
	var vegAmount float64
	fmt.Println("Adding a new vegetable ")
	fmt.Println("Enter vegetable Name : ")
	_, _ = fmt.Scanln(&vegName)
	fmt.Println("Enter vegetable Price : ")
	_, _ = fmt.Scanln(&vegPrice)
	fmt.Println("Enter vegetable Amount : ")
	_, _ = fmt.Scanln(&vegAmount)

	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`
	var veg common.Vegetable
	if err := client.Call("Vegetables.Add", common.Vegetable{
		Name:   vegName,
		Price:  vegPrice,
		Amount: vegAmount,
	}, &veg); err != nil {
		fmt.Println("Error:2 Vegetables.Add()", err)
	} else {
		fmt.Printf("Success: Added '%s' \n", veg.Name)
	}
	fmt.Println("---------------------- End of Operation -------------------------")
}

func updateVegetable() {
	var vegName string
	var vegPrice float64
	var vegAmount float64
	fmt.Println("Updating a vegetable ")
	fmt.Println("Enter vegetable Name : ")
	_, _ = fmt.Scanln(&vegName)
	fmt.Println("Enter new vegetable Price : ")
	_, _ = fmt.Scanln(&vegPrice)
	fmt.Println("Enter new vegetable Amount : ")
	_, _ = fmt.Scanln(&vegAmount)

	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`
	var veg common.Vegetable
	if err := client.Call("Vegetables.Update", common.Vegetable{
		Name:   vegName,
		Price:  vegPrice,
		Amount: vegAmount,
	}, &veg); err != nil {
		fmt.Println("Error:2 Vegetables.Update()", err)
	} else {
		fmt.Printf("Success: Updated '%s' \n", veg.Name)
	}
	fmt.Println("---------------------- End of Operation -------------------------")
}
