package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Test represents a user with validation rules
type Test struct {
	Dir string `validate:"dirpath"`
}

func main() {
	validate := validator.New()

	user := Test{
		Dir: "/tmp/bananas",
	}

	// Validate the struct
	err := validate.Struct(user)
	if err != nil {
		// Handling validation errors
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation with tag '%s'\n", err.Field(), err.Tag())
		}
	} else {
		fmt.Println("All fields are valid!")
	}
}
