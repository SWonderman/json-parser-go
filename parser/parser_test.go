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
