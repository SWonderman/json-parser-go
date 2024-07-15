# Simple JSON Parser written in Go

### Introduction
As a way of getting to know Go a little bit more, I decided to create a very simple JSON parser.

The idea of the program is more simple than its implementation. You have a stringified JSON file, you pass it to the parser, and assuming that the input is valid JSON, you get a formatted Go 'map' as an output that allows you to easily access different input values.

The parser works with single objects, arrays of objects, integers, floats, and nested elements. If the JSON input is invalid, the location of the error with an appropriate error message will be returned to the user.

A little note, even though the trailing comma is not a valid JSON, which technically should be reported to the user, the parser will omit the trailing comma and parse the input without complaining.

The general design/structure of the parser was inspired by ["Writing An Interpreter In Go"](https://interpreterbook.com/) by Thorsten Ball book.


### Usage
The parsing is done by calling a facade function `Parse` with a string input from the `jsonparser.go` file. The function will return a parser result and parser errors. 

Since the JSON can either be a single object or an array of objects, the parser result will either have a `map[string]any` or a `[]map[string]any]` types. The parser result comes with two handy methods for checking if a single map or an array of maps was returned as a result - `IsSingleMap()` and `IsMapArray()`. After knowing the type of the parsing result, the result values can be accessed by accessing the appropriate member variable`result.SingleMap`or `result.MapArray`.

The parser errors is just an array of strings with meaningful error messages showing where and why it was not possible to produce a valid result.


Here is a basic usage:
```go
import "sw/json-parser/jsonparser"

func main() {
    input := `{"first_name": "Joe", "last_name": "Doe", "age": 88}`
    result, errors := jsonparser.Parse(input)
    if errors != nil {
        for _, err := range errors {
            fmt.Println(err)
        }
        return
    }

    /*
    The for-loop will produce:
        first_name -> Joe
        last_name -> Doe
        age -> %!s(int=88)
    */
    if result.IsSingleMap() {
        for key, value := range result.SingleMap {
            fmt.Printf("%s -> %s\n", key, value)
        }
    }
}
```

As mentioned, if the input is not a valid JSON, an errors response will show what is wrong. For example, for the following input
```go
input := `{first_name: "Joe", "last_name": "Doe", "age": 88}
```
the resulting error will say

    PARSER ERROR: line 1 and column 2 near token literal 'first_name'.
    Key value has to be of type string. Did you add quotation marks around the key value?

and for this faulty input
```go
input := `{"first_name": "Joe", "last_name": "Doe", "age": 88.88.88}`
```
the error will say

    PARSER ERROR: line 1 and column 50 near token literal '88.88.88'.
    It was not possible to parse the token literal as either int or float.

### More Examples
Parsing an array of objects
```go
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
```

A bit more complex object with a nested array and an infamous trailing comma
```go
input := `{"name": "Joe", "age": 7, "orders": [{"name": "foo", "price": 11.99},]}`
result, errors := jsonparser.Parse(input)
if errors != nil {
    for _, err := range errors {
        fmt.Println(err)
    }
    return
}

if result.IsSingleMap() {
    if result.IsSingleMap() {
        for key, value := range result.SingleMap {
            fmt.Printf("%s -> %s\n", key, value)
        }
    }
}
```