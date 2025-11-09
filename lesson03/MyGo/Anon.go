package MyGo

//包导入相对于模块的根目录
import (
	"MyL3/ty"
	"fmt"
)

type Anon struct {
	Info ty.Member
}

type AnonDo interface {
	SelfIntroduction()
	TangLaugh()
	C_Chord()
}

func (a *Anon) SelfIntroduction() {
	fmt.Println("Chihaya Anon desu!!")
}

func (a *Anon) TangLaugh() {
	fmt.Println("eiheiheihei...")
}

func (a *Anon) C_Chord() {
	fmt.Println("C ko-do!")
}
