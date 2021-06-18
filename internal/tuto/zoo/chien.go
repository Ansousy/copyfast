package zoo

import "fmt"

type Chien struct {
	Pattes int
	word   string
	Name   string
}

func (c *Chien) Say() {
	fmt.Println(c.word)
}

func MakeChien(Name string, word string) *Chien {

	return &Chien{
		Pattes: 4,
		Name:   Name,
		word:   word,
	}
}
