package parser

import (
	"testing"

	"sw/json-parser/lexer"
)

func TestParserSimpleString(t *testing.T) {
	input := `{"name": "Joe"}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]string{
		"name": "Joe",
	}

	for key, value := range expectedMap {
		if parserResult.SingleMap[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parserResult.SingleMap[key])
		}
	}
}

func TestParserSimpleInt(t *testing.T) {
	input := `{"age": 88}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]int{
		"age": 88,
	}

	for key, value := range expectedMap {
		if parserResult.SingleMap[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%d, but got %q->%d", key, value, key, parserResult.SingleMap[key])
		}
	}
}

func TestParserNegativeInt(t *testing.T) {
	input := `{"age": -88}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]int{
		"age": -88,
	}

	for key, value := range expectedMap {
		if parserResult.SingleMap[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%d, but got %q->%d", key, value, key, parserResult.SingleMap[key])
		}
	}
}

func TestParserSimpleFloat(t *testing.T) {
	input := `{"salary": 99.78}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]float64{
		"salary": 99.78,
	}

	for key, value := range expectedMap {
		if parserResult.SingleMap[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%f, but got %q->%f", key, value, key, parserResult.SingleMap[key])
		}
	}
}

func TestParserSimpleStringOnlyObject(t *testing.T) {
	input := `{"first_name": "Joe", "last_name": "Doe"}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]string{
		"first_name": "Joe",
		"last_name":  "Doe",
	}

	for key, value := range expectedMap {
		if parserResult.SingleMap[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parserResult.SingleMap[key])
		}
	}
}

func TestParserSimpleIntOnlyObject(t *testing.T) {
	input := `{"age": 88, "salary": 1000}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]int{
		"age":    88,
		"salary": 1000,
	}

	for key, value := range expectedMap {
		if parserResult.SingleMap[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parserResult.SingleMap[key])
		}
	}
}

func TestParserSimpleStringArray(t *testing.T) {
	input := `{"colors": ["blue", "red", "green"]}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string][]string{
		"colors": []string{"blue", "red", "green"},
	}

	for key, value := range expectedMap {
		colors := value
		parsedColors := parserResult.SingleMap[key]

		if slice, ok := parsedColors.([]string); ok {
			for idx, color := range colors {
				if color != slice[idx] {
					t.Fatalf("Parsed value in the array does not match with what is expected. Got %q, but expected %q", slice[idx], color)
				}
			}
		}
	}
}

func TestParserKeywords(t *testing.T) {
	input := `{"hasKids": false, "hasJob": true, "ownsCar": null}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]any{
		"hasKids": false,
		"hasJob":  true,
		"ownsCar": nil,
	}

	for key, value := range expectedMap {
		if parserResult.SingleMap[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parserResult.SingleMap[key])
		}
	}
}

func TestParserSimpleObject(t *testing.T) {
	input := `{"name": "Joe", "age": 88, "salary": 99.78, "colors": ["blue", "green", "red"], "hasKids": false, "hasJob": true, "ownsCar": null}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]any{
		"name":    "Joe",
		"age":     88,
		"salary":  99.78,
		"colors":  []string{"blue", "green", "red"},
		"hasKids": false,
		"hasJob":  true,
		"ownsCar": nil,
	}

	for key, value := range expectedMap {
		if _, ok := parserResult.SingleMap[key].([]any); ok {
			// skip checking values that are arrays
		} else {
			if parserResult.SingleMap[key] != value {
				t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parserResult.SingleMap[key])
			}
		}
	}

	parsedColors, ok := parserResult.SingleMap["colors"].([]any)
	if ok == false {
		t.Fatalf("Colors array was not parsed and is missing in the output")
	}

	expectedColors, ok := expectedMap["colors"].([]string)
	if ok == false {
		t.Fatalf("This will not happen...")
	}

	for idx, color := range expectedColors {
		if color != parsedColors[idx] {
			t.Fatalf("Parsed colors do not match. Expected %q, but got %q", color, parsedColors[idx])
		}
	}
}

func TestParserArrayOfObjects(t *testing.T) {
	input := `{"orders": [{"id": 12, "article": "book"}, {"id": 13, "article": "ball"}]}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parserResult, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

    if parserResult.IsSingleMap() == false {
        t.Fatalf("Parser result is not a single map")
    }

	expectedMap := map[string]any{
		"orders": []map[string]any{
			{
				"id":      12,
				"article": "book",
			},
			{
				"id":      13,
				"article": "ball",
			},
		},
	}

	parsedOrders, ok := parserResult.SingleMap["orders"].([]any)
	if ok == false {
		t.Fatalf("Orders data was not parsed")
	}

	expectedOrders, ok := expectedMap["orders"].([]map[string]any)
	if ok == false {
		t.Fatalf("This will not happen...")
	}

	for idx, expected := range expectedOrders {
		parsedOrder, ok := parsedOrders[idx].(map[string]any)
		if ok == false {
			t.Fatalf("Parsed order is not a map")
		}

		for key, value := range expected {
			if parsedOrder[key] != value {
				t.Fatalf("Parsed nested object does not match with the expected object.")
			}
		}
	}
}
