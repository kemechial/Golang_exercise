package main

import (
	"errors"
	"fmt"
)

// User struct
type User struct {
	Name string
	Age  int
}

// Constructor returning struct
func NewUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

// Constructor returning pointer with validation
func NewUserWithValidation(name string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if age < 0 {
		return nil, errors.New("age cannot be negative")
	}
	return &User{
		Name: name,
		Age:  age,
	}, nil
}

// Another example returning a pointer explicitly
func NewUserPointer(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func main() {
	// Example 1: Struct returned
	user1 := NewUser("Kaan", 25)
	fmt.Println("Struct returned:", user1)

	// Example 2: Pointer with validation
	user2, err := NewUserWithValidation("Kaan", 25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pointer with validation:", *user2)
	}

	// Example 3: Pointer returned directly
	user3 := NewUserPointer("Kaan", 25)
	fmt.Println("Pointer returned directly:", *user3)
}