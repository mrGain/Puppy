package puppy

import (
	"github.com/mrGain/dmgmt_dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}