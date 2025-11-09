package MyGo

import (
	"MyL3/ty"
	"fmt"
)

type Rāna struct {
	Info ty.Member
}

type RānaDo interface {
	PlayHaruhikage()
	Interesting()
}

func (r *Rāna) PlayHaruhikage() {
	fmt.Println("(music)32123432~")
}

func (r *Rāna) Interesting() {
	fmt.Println("Interesting girl~")
}
