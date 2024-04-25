package goaa

import (
	"encoding/json"
	"log"
	"os"
)

func Test() {
	test1()
}
func test1() {
	from := "miami"
	to := "texas"
	departDate := "2024-05-01"
	returnDate := "2024-05-04"
	passengers := 1
	locale := "es_EC"
	airports1, err := Airports(from, nil)
	if err != nil {
		log.Println("test -> 1 -> err: ", err)
		return
	}
	airports2, err := Airports(to, nil)
	if err != nil {
		log.Println("test -> 2 -> err: ", err)
		return
	}
	f1, _ := os.Create("./suggestions1.json")
	json.NewEncoder(f1).Encode(airports1)
	f2, _ := os.Create("./suggestions2.json")
	json.NewEncoder(f2).Encode(airports2)
	results, err := Flights(locale, airports1[0].Code, airports2[1].Code, departDate, returnDate, passengers, nil)
	if err != nil {
		log.Println("test -> 3 -> err: ", err)
		return
	}
	f3, _ := os.Create("./results.json")
	json.NewEncoder(f3).Encode(results)
}
