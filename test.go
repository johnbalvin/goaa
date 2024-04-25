package goaa

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	Test()
}
func Test() {
	test1()
}
func test1() {
	from := "new york"
	to := "galapagos"
	departDate := "2024-05-01"
	returnDate := "2024-05-04"
	passengers := 1
	locale := "es_EC" //where you are located, probably for increasing the price or is just for statistics, I DON'T KNOW, do not say that I said this field is for incresing the price, it's jut a theory
	airports1, err := Airports(from, nil)
	if err != nil {
		log.Println(err)
		return
	}
	airports2, err := Airports(to, nil)
	if err != nil {
		log.Println(err)
		return
	}
	flights, err := Flights(locale, airports1[0].Code, airports2[0].Code, departDate, returnDate, passengers, nil)
	if err != nil {
		log.Println(err)
		return
	}
	rawJSON, _ := json.MarshalIndent(flights, "", "  ")
	fmt.Printf("%s", rawJSON) //in case you don't have write permisions
	if err := os.WriteFile("./flights.json", rawJSON, 06444); err != nil {
		log.Println(err)
		return
	}
}
