package main

import (
    "fmt"

    "sw/json-parser/jsonparser"
)

func main() {
    input := `{"name": "Joe"}`
    parserResult, error := jsonparser.Parse(input)
    if error != nil {
        fmt.Println(error.Error())
    }

    if parserResult.IsSingleMap() {
        fmt.Printf("For 'name' key got: %s\n", parserResult.SingleMap["name"])
    }
}
