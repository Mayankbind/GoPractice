package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Try changing the JSON below to practice
	jsonSamples := []string{
		`{"name": "Mayank", "age": 26, "active": true}`,
		`["go", "python", 123, false]`,
		`"hello world"`,
		`12345`,
		`true`,
	}

	for _, sample := range jsonSamples {
		fmt.Println("\n--- Parsing JSON ---")
		fmt.Println(sample)
		
		
		var any interface{}
		err := json.Unmarshal([]byte(sample), &any)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Result Type Switch:")
		handle(any)
	}
}

// type dum interface {
// 	lol() error
// }


// type tt struct {}

// func (t *tt) lol() error {
// 	return nil
// }

// var _ dum = (*tt)(nil)

func handle(v interface{}) {
	switch t := v.(type) {
	case map[string]interface{}:
		fmt.Println("Type: Object (map)")
		for key, value := range t {
			fmt.Printf("  Key: %s → ", key)
			handle(value)
		}

	case []interface{}:
		fmt.Println("Type: Array")
		for i, value := range t {
			fmt.Printf("  Index %d → ", i)
			handle(value)
		}

	case string:
		fmt.Println("Type: String →", t)

	case float64:
		fmt.Println("Type: Number →", t)

	case bool:
		fmt.Println("Type: Boolean →", t)

	case nil:
		fmt.Println("Type: Null")

	default:
		fmt.Println("Unknown type:", t)
	}
}
