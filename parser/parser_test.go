package parser

import (
	"testing"

	"sw/json-parser/lexer"
)

func TestParserSimpleString(t *testing.T) {
	input := `{"name": "Joe"}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

	expectedMap := map[string]string{
		"name": "Joe",
	}

	for key, value := range expectedMap {
		if parsed[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parsed[key])
		}
	}
}

func TestParserSimpleNumber(t *testing.T) {
	input := `{"age": 88}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

	expectedMap := map[string]int{
		"age": 88,
	}

	for key, value := range expectedMap {
		if parsed[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%d, but got %q->%d", key, value, key, parsed[key])
		}
	}
}

func TestParserSimpleStringOnlyObject(t *testing.T) {
	input := `{"first_name": "Joe", "last_name": "Doe"}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

	expectedMap := map[string]string{
		"first_name": "Joe",
		"last_name":  "Doe",
	}

	for key, value := range expectedMap {
		if parsed[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parsed[key])
		}
	}
}

func TestParserSimpleIntOnlyObject(t *testing.T) {
	input := `{"age": 88, "salary": 1000}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

	expectedMap := map[string]int{
		"age":    88,
		"salary": 1000,
	}

	for key, value := range expectedMap {
		if parsed[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parsed[key])
		}
	}
}

func TestParserSimpleStringArray(t *testing.T) {
	input := `{"colors": ["blue", "red", "green"]}`

	lexer := lexer.New(input)
	parser := New(lexer)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

	expectedMap := map[string][]string{
		"colors": []string{"blue", "red", "green"},
	}

	for key, value := range expectedMap {
		colors := value
		parsedColors := parsed[key]

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

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parser returned an error. Error: %q", err)
	}

	expectedMap := map[string]any{
		"hasKids": false,
		"hasJob":  true,
		"ownsCar": nil,
	}

	for key, value := range expectedMap {
		if parsed[key] != value {
			t.Fatalf("Parser returned an unexpected key-value pair. Expected: %q->%q, but got %q->%q", key, value, key, parsed[key])
		}
	}
}
