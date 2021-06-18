package tuto

import (
	"fmt"

	"github.com/Devopsengineer75/copyfastgo/internal/tuto/zoo"
)

func If() bool {

	a := 0

	if a > 0 {
		return true
	}

	return false
}

func IfElse() bool {
	a := 0

	if a > 0 {
		return true
	} else {
		return false
	}
}

func IfElseIfElse() string {
	a := 0

	if a == 0 {
		return "equal"
	} else if a > 1 {
		return "superior"
	} else {
		return "inferior"
	}
}

func SwitchCase(animal zoo.IsAnimal) {
	switch a := animal.(type) {
	case *zoo.Chat:
		fmt.Println("Super je peux adopter", a.Name)
	case *zoo.Chien:
		fmt.Println("Super je peux adopter", a.Name)
	case *zoo.Oiseau:
		fmt.Println("Super je peux adopter", a.Name)
	}
}
