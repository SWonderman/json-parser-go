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
			result, ok := parsedMap.(map[string]any)
			if ok == false {
				fmt.Println("Parser result is not a map when it was expected to be!")
			}

			for key, value := range result {
				fmt.Printf("[%d] %s->%s\n", idx, key, value)
			}
		}
	}
}
