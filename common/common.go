package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Vegetable struct represents a vegetable.
type Vegetable struct {
	Name          string
	Price, Amount float64
}

// Vegetables struct represents a list of Vegetables.
type Vegetables struct {
	vegetables map[string]Vegetable // private
}

// Add a new vegetable to the file.
func (c *Vegetables) Add(veg Vegetable, reply *Vegetable) error {

	// Check if the vegetable already exists in the file.
	if _, ok := c.vegetables[veg.Name]; ok {
		return fmt.Errorf("vegetable with name '%s' already exists", veg.Name)
	}

	c.vegetables[veg.Name] = veg

	// Append the new vegetable to the file.
	file, err := os.OpenFile("veg-data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	price := fmt.Sprintf("%0.0f", veg.Price)
	amt := fmt.Sprintf("%0.0f", veg.Amount)
	_, e := file.WriteString(veg.Name + "," + price + "," + amt + "\n")
	if e != nil {
		fmt.Println(err)
	}

	file.Sync()
	// set reply value
	*reply = veg

	// return `nil` error
	return nil
}

// Update the details of a vegetable.
func (c *Vegetables) Update(veg Vegetable, updatedItem *Vegetable) error {

	// Check if vegetable already exists in the file.
	if _, ok := c.vegetables[veg.Name]; !ok {
		return fmt.Errorf("vegetable with name '%s' doesn't exists", veg.Name)
	}

	if veg.Price == 0 {
		veg.Price = c.vegetables[veg.Name].Price
	}
	if veg.Amount == 0 {
		veg.Amount = c.vegetables[veg.Name].Amount
	}

	c.vegetables[veg.Name] = veg
	WriteFile(c)

	// set reply value
	*updatedItem = veg

	// return `nil` error
	return nil
}

// Returns a vegetable with specific name (procedure).
func (c *Vegetables) Get(name string, reply *Vegetable) error {

	// get vegetable with name from the file
	result, ok := c.vegetables[name]

	// check if vegetable exists in the file
	if !ok {
		return fmt.Errorf("vegetable with name '%s' does not exist", name)
	}

	// set reply value
	*reply = result

	// return `nil` error
	return nil
}

//Returns the vegetable list.
func (c *Vegetables) GetAll(payload string, vegItems *[]string) error {

	names := make([]string, 0, len(c.vegetables))
	for key := range c.vegetables {
		names = append(names, key)
	}

	// set vegetable List
	*vegItems = names

	// return `nil` error
	return nil
}

// Returns a new instance of Vegetables (pointer).
func ReadFile() *Vegetables {
	fmt.Println("Read data from file")
	f, err := os.Open("veg-data.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// Loop through lines & turn into object
	wholeList := make(map[string]Vegetable)
	for scanner.Scan() {
		nxt := strings.Split(scanner.Text(), ",")
		price, _ := strconv.ParseFloat(nxt[1], 64)
		amount, _ := strconv.ParseFloat(nxt[2], 64)
		wholeList[nxt[0]] = Vegetable{
			Name:   nxt[0],
			Price:  price,
			Amount: amount,
		}
	}
	defer f.Close()
	fmt.Println("Successfully read data")

	return &Vegetables{
		vegetables: wholeList,
	}

}

// Write data to the file.
func WriteFile(veg *Vegetables) {
	fmt.Println("Write data to file")
	f, err := os.OpenFile("veg-data.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for item := range veg.vegetables {
		name := fmt.Sprintf("%s", veg.vegetables[item].Name)
		price := fmt.Sprintf("%0.0f", veg.vegetables[item].Price)
		amount := fmt.Sprintf("%0.0f", veg.vegetables[item].Amount)

		_, err := f.WriteString(name + "," + price + "," + amount + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}

	f.Sync()

}
