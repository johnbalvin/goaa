package airports

import (
	"encoding/json"
	"log"
	"os"
)

func Test() {
	test1()
}
func test1() {
	searchText := "miami"
	data, err := Get(searchText, nil)
	if err != nil {
		log.Println("test1 -> err: ", err)
		return
	}
	f, _ := os.Create("./suggestions.json")
	json.NewEncoder(f).Encode(data)
}
