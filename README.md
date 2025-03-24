# Go Data Types REST API

This educational project is a simple REST API created to learn about data types in the Go programming language. The API demonstrates basic Go data types with examples.

## Overview

The API provides the most fundamental and commonly used data types in Go through a single endpoint:

- Numeric Types (int, float64)
- Text Types (string)
- Boolean Type (bool)
- Array Types (array, slice)
- Map Type
- Struct Type

## API Endpoint
### GET /go-data-types

Returns all Go data types in JSON format.

#### Example Request
```bash
curl http://localhost:8080/go-data-types
```

#### Example Response
```json
{
  "int": {
    "description": "Basic integer type",
    "example": 42,
    "usage": "var age int = 42"
  },
  "string": {
    "description": "Text type",
    "example": "Hello World!",
    "usage": "var message string = \"Hello World!\""
  },
  "slice": {
    "description": "Dynamic-size array",
    "example": ["apple", "banana"],
    "usage": "var fruits []string = []string{\"apple\", \"banana\"}"
  }
}
```

## Installation and Running

1. Clone the project
2. Navigate to project directory
3. Start the server:
```bash
go run main.go
```

The server will run on port 8080 by default.

## About Data Types

This API demonstrates the most basic data types in the Go programming language. For each data type, it provides:
- Simple and clear descriptions
- Usage examples

This makes it easier for beginners to understand data types in Go.


## Contact

A.Samet Palitci - [@asametpalitci](https://twitter.com/asametpalitci) - abdulsametpalitci@gmail.com

Project Link: [https://github.com/sametpalitci/go-data-types](https://github.com/sametpalitci/go-data-types)