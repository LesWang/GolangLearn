package main

import "fmt"

func main()  {
    var countryCapitalMap map[string] string
    countryCapitalMap = make(map[string] string)

    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"
    countryCapitalMap["Japan"] = "Tokyo"

    for country := range countryCapitalMap {
        fmt.Printf("%s's capital is %s\n", country, countryCapitalMap[country])
    }

    fmt.Println("delete France")
    delete(countryCapitalMap, "France")
    for country := range countryCapitalMap {
        fmt.Printf("%s's capital is %s\n", country, countryCapitalMap[country])
    }

    capital, ok := countryCapitalMap["America"]
    fmt.Println(capital)
    fmt.Println(ok)
    capital2, ok2 := countryCapitalMap["Italy"]
    fmt.Println(capital2)
    fmt.Println(ok2)
}
