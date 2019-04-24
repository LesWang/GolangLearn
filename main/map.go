package main

import "fmt"

func mapTest()  {
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

func wordCount(s string) map[string]int {
    var wordMap = make(map[string]int)
    var start, end = 0, 0
    for index, ch := range s {
        if ch == ' ' {
            end = index
            if start != end {
                addWordRecord(s[start: end], wordMap)
            }
            start = end + 1
        }
    }
    if start < len(s) - 1 {
        addWordRecord(s[start:], wordMap)
    }
    return wordMap
}

func addWordRecord(word string, wordMap map[string]int)  {
    if _, ok := wordMap[word]; ok {
        wordMap[word]++
    } else {
        wordMap[word] = 1
    }
}

func main()  {
    fmt.Println(wordCount("Hello world yes ok    i i i "))
}
