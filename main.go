package main

import (
	"fmt"

	"sw/json-parser/jsonparser"
)

func main() {
	input := `[{"name": "Joe"}, {"name": "Kevin"}]`
	parserResult, error := jsonparser.Parse(input)
	if error != nil {
		fmt.Println(error.Error())
	}

	if parserResult.IsMapArray() {
		for idx, parsedMap := range parserResult.MapArray {
			for key, value := range parsedMap{
				fmt.Printf("[%d] %s->%s\n", idx, key, value)
			}
		}
	}
}
