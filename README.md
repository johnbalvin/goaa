# American Airlines scraper in Go

## Overview
This project is an open-source tool developed in Golang for extracting product information from American Airlines. It's designed to be fast, efficient, and easy to use, making it an ideal solution for developers looking for American Airlines flights.

## Features
- Full search support
- Extracts detailed product information from American Airlines
- Implemented in Go for performance and efficiency
- Easy to integrate with existing Go projects

## Examples

### Getting airports

```Go
    package main

    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
        "github.com/johnbalvin/goaa"
    )
    func main(){
        from := "miami"
        to:=    "texas"
        airports1, err := goaa.Airports(from, nil)
        if err != nil {
            log.Println(err)
            return
        }
        airports2, err := goaa.Airports(to, nil)
        if err != nil {
            log.Println(err)
            return
        }
        rawJSON1, _ := json.MarshalIndent(airports1, "", "  ")
        fmt.Printf("%s", rawJSON1) //in case you don't have write permisions
        if err := os.WriteFile("./airports1.json", rawJSON1, 06444); err != nil {
            log.Println(err)
            return
        }
        rawJSON2, _ := json.MarshalIndent(airports2, "", "  ")
        fmt.Printf("%s", rawJSON2) //in case you don't have write permisions
        if err := os.WriteFile("./airports2.json", rawJSON2, 06444); err != nil {
            log.Println(err)
            return
        }
    }
```

### Getting flights
```Go
    package main

    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
        "github.com/johnbalvin/goaa"
    )
    func main(){
        originAirport := "GYE"
        destinationAirport := "MIA"
        departDate := "2024-05-01"
        returnDate := "2024-05-04"
        passengers := 1
        locale := "es_EC"//where you are located, probably for increasing the price or is just for statistics, I DON'T KNOW, do not say that I said this field is for incresing the price, it's jut a theory
        flights, err := goaa.Flights(locale, originAirport, destinationAirport, departDate, returnDate, passengers, nil)
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
```

### Usage combine

```Go
    package main

    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
        "github.com/johnbalvin/goaa"
    )
    func main(){
        from := "new york"
        to:=    "galapagos"
        departDate := "2024-05-01"
        returnDate := "2024-05-04"
        passengers := 1
        locale := "es_EC"//where you are located, probably for increasing the price or is just for statistics, I DON'T KNOW, do not say that I said this field is for incresing the price, it's jut a theory
        airports1, err := goaa.Airports(from, nil)
        if err != nil {
            log.Println(err)
            return
        }
        airports2, err := goaa.Airports(to, nil)
        if err != nil {
            log.Println(err)
            return
        }
        flights, err := goaa.Flights(locale, airports1[0].Code, airports2[0].Code, departDate, returnDate, passengers, nil)
        if err != nil {
            log.Println(err)
            return
        }
    }
```