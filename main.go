package main

import (
	"fmt"

	"sw/json-parser/jsonparser"
)

func parseArrayOfObjects() {
	input := `[{"name": "Joe"}, {"name": "Kevin"}]`
	parserResult, errors := jsonparser.Parse(input)
	if errors != nil {
        for _, err := range errors {
            fmt.Println(err)
        }
        return
	}

	if parserResult.IsMapArray() {
		for idx, parsedMap := range parserResult.MapArray {
			for key, value := range parsedMap {
				fmt.Printf("[%d] %s->%s\n", idx, key, value)
			}
		}
	}
}

func main() {
    input := `{"name": "Joe", "age": -7, "orders": [{"name": "foo", "price": 11.99},]}`
    result, errors := jsonparser.Parse(input)
	if errors != nil {
        for _, err := range errors {
            fmt.Println(err)
        }
        return
	}

    if result.IsSingleMap() {
        fmt.Println(result.SingleMap["orders"])
    }
}
