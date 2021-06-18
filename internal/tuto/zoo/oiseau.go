package zoo

import "fmt"

type Oiseau struct {
	Pattes int
	word   string
	Name   string
}

func (c *Oiseau) Say() {
	fmt.Println(c.word)
}

func MakeOiseau(Name string, word string) *Oiseau {

	return &Oiseau{
		Pattes: 2,
		Name:   Name,
		word:   word,
	}
}
