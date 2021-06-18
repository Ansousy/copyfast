package zoo

import "fmt"

type Chat struct {
	Pattes int
	word   string
	Name   string
}

func (c *Chat) Say() {
	fmt.Println(c.word)
}

func MakeChat(Name string, word string) *Chat {

	return &Chat{
		Pattes: 4,
		Name:   Name,
		word:   word,
	}
}
