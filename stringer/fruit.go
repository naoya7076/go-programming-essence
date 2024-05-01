package main

import "fmt"

//go:generate stringer -type Fruit fruit.go
type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func (i Fruit) String() string {
	switch i {
	case Apple:
		return "Apple"
	case Orange:
		return "Orange"
	case Banana:
		return "Banana"
	}
	return fmt.Sprintf("Fruit(%d)", i)
}
