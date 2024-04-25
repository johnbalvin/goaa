package flights

import (
	"encoding/json"
	"log"
	"os"
)

func Test() {
	test1()
}
func test1() {
	originAirport := "GYE"
	destinationAirport := "MIA"
	departDate := "2024-05-01"
	returnDate := "2024-05-04"
	passengers := 1
	locale := "es_EC"
	data, err := Get(locale, originAirport, destinationAirport, departDate, returnDate, passengers, nil)
	if err != nil {
		log.Println("test1 -> err: ", err)
		return
	}
	f, _ := os.Create("./data.json")
	json.NewEncoder(f).Encode(data)
}
