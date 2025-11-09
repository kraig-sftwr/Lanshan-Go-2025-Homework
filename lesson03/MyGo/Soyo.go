package MyGo

import (
	"MyL3/ty"
	"fmt"
)

type Soyo struct {
	Info ty.Member
}

type SoyoDo interface {
	Roar()
	Plead()
}

func (s *Soyo) Roar() {
	fmt.Println("Why played Haruhikage!!!")
}

func (s *Soyo) Plead() {
	fmt.Println("(kneel)O ne ga i!")
}
