package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type DataTypeExample struct {
	Description string      `json:"description"`
	Example     interface{} `json:"example"`
	Usage       string      `json:"usage"`
}

func handleGoDataTypes(w http.ResponseWriter, r *http.Request) {
	dataTypes := map[string]DataTypeExample{
		// Numeric
		"int": {
			Description: "Signed integer type (platform dependent size)",
			Example:     42,
			Usage:       "var age int = 42",
		},
		"int8": {
			Description: "8-bit signed integer (-128 to 127)",
			Example:     127,
			Usage:       "var temperature int8 = 127",
		},
		"int16": {
			Description: "16-bit signed integer (-32768 to 32767)",
			Example:     32767,
			Usage:       "var year int16 = 2024",
		},
		"int32": {
			Description: "32-bit signed integer (-2147483648 to 2147483647)",
			Example:     2147483647,
			Usage:       "var population int32 = 2147483647",
		},
		"int64": {
			Description: "64-bit signed integer",
			Example:     9223372036854775807,
			Usage:       "var distance int64 = 9223372036854775807",
		},
		"uint": {
			Description: "Unsigned integer type (platform dependent size)",
			Example:     42,
			Usage:       "var count uint = 42",
		},
		"uint8": {
			Description: "8-bit unsigned integer (0 to 255)",
			Example:     255,
			Usage:       "var brightness uint8 = 255",
		},
		"uint16": {
			Description: "16-bit unsigned integer (0 to 65535)",
			Example:     65535,
			Usage:       "var port uint16 = 65535",
		},
		"uint32": {
			Description: "32-bit unsigned integer (0 to 4294967295)",
			Example:     4294967295,
			Usage:       "var fileSize uint32 = 4294967295",
		},
		"uint64": {
			Description: "64-bit unsigned integer",
			Example:     1844674407370955161,
			Usage:       "var memory uint64 = 1844674407370955161",
		},
		"float32": {
			Description: "32-bit floating point number",
			Example:     3.14159,
			Usage:       "var pi float32 = 3.14159",
		},
		"float64": {
			Description: "64-bit floating point number",
			Example:     3.14159265359,
			Usage:       "var pi float64 = 3.14159265359",
		},

		// Text
		"string": {
			Description: "UTF-8 encoded text",
			Example:     "Hello World",
			Usage:       "var message string = \"Hello World\"",
		},
		"rune": {
			Description: "UTF-8 character (alias for int32)",
			Example:     "A",
			Usage:       "var letter rune = 'A'",
		},
		"byte": {
			Description: "Alias for uint8",
			Example:     255,
			Usage:       "var b byte = 255",
		},

		// Boolean
		"bool": {
			Description: "true or false",
			Example:     true,
			Usage:       "var isActive bool = true",
		},

		// Complex
		"complex64": {
			Description: "Complex number with float32 real and imaginary parts",
			Example:     "3+4i",
			Usage:       "var c complex64 = 3 + 4i",
		},
		"complex128": {
			Description: "Complex number with float64 real and imaginary parts",
			Example:     "3.14+4.2i",
			Usage:       "var c complex128 = 3.14 + 4.2i",
		},

		// Composite
		"array": {
			Description: "Fixed-size sequence of elements",
			Example:     []int{1, 2, 3, 4, 5},
			Usage:       "var numbers [5]int = [5]int{1, 2, 3, 4, 5}",
		},
		"slice": {
			Description: "Dynamic-size sequence of elements",
			Example:     []string{"elma", "armut", "muz"},
			Usage:       "var fruits []string = []string{\"elma\", \"armut\", \"muz\"}",
		},
		"map": {
			Description: "Key-value pairs collection",
			Example:     map[string]int{"bir": 1, "iki": 2, "üç": 3},
			Usage:       "var numbers map[string]int = map[string]int{\"bir\": 1, \"iki\": 2, \"üç\": 3}",
		},
		"struct": {
			Description: "Collection of fields",
			Example: struct {
				Name string
				Age  int
			}{"Samet", 30},
			Usage: "type Person struct { Name string; Age int }",
		},
		"pointer": {
			Description: "Points to memory address of another value",
			Example:     "&42",
			Usage:       "var ptr *int = &value",
		},

		// Special
		"interface": {
			Description: "Set of method signatures",
			Example:     "interface{}",
			Usage:       "var any interface{} = 42",
		},
		"channel": {
			Description: "Communication between goroutines?",
			Example:     "chan int",
			Usage:       "ch := make(chan int)",
		},
		"uintptr": {
			Description: "Unsigned integer to store pointer values",
			Example:     123456789,
			Usage:       "var ptr uintptr = uintptr(unsafe.Pointer(&value))",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataTypes)
}

func main() {
	http.HandleFunc("/go-data-types", handleGoDataTypes)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
