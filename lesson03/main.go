package main

import (
	"MyL3/MyGo"
	"MyL3/ty"
)

func main() {
	anon := MyGo.Anon{
		Info: ty.Member{
			Name:     "Chihaya Anon",
			Position: "Rhythm Guitar",
		},
	}

	soyo := MyGo.Soyo{
		Info: ty.Member{
			Name:     "Nagasaki Soyo",
			Position: "Bass",
		},
	}

	rāna := MyGo.Rāna{
		Info: ty.Member{
			Name:     "Kaname Rāna",
			Position: "Lead Guitar",
		},
	}

	anon.SelfIntroduction()
	rāna.PlayHaruhikage()
	soyo.Roar()

}
