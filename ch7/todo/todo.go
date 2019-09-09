package main

import "time"

// Todo include four attr
type Todo struct {
	ID       int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is the slice of Todo
type Todos []Todo
