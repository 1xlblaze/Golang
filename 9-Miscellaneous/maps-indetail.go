package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Creating a map with mixed types
	data := map[string]interface{}{
		"name":   "Alice",
		"age":    25,   // Int type
		"height": 5.7,  // Float type
		"active": true, // Boolean type
		"scores": []int{90, 85, 88},
		"meta": map[string]string{
			"role": "admin",
		},
	}

	// Safe retrieval without panic
	if value, ok := data["age"]; ok {
		if age, ok := value.(int); ok {
			fmt.Println("Age:", age)
		} else {
			fmt.Println("Error: 'age' is not an int")
		}
	}

	// Safe retrieval with JSON example
	jsonStr := `{"name": "Bob", "age": 30, "height": 6.1, "active": false}`
	var jsonData map[string]interface{}

	// Unmarshal JSON into map
	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Handling numbers (JSON defaults to float64)
	if value, ok := jsonData["age"]; ok {
		if age, ok := value.(float64); ok {
			fmt.Println("Age (converted to int):", int(age))
		}
	}

	// Using a type switch for dynamic values
	fmt.Println("\nIterating over map with type detection:")
	for key, value := range jsonData {
		switch v := value.(type) {
		case string:
			fmt.Printf("%s is a string: %s\n", key, v)
		case float64:
			fmt.Printf("%s is a number: %.2f\n", key, v)
		case bool:
			fmt.Printf("%s is a boolean: %t\n", key, v)
		default:
			fmt.Printf("%s is of unknown type\n", key)
		}
	}
}
