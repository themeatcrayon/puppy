package puppy

/*
Adding a tag for v1.0.0
*/

import (
	"github.com/prinsb/dog"
)

func Bark() string {
	return "Bark!"
}
func Barks() string {
	return "bark bark bark!"
}

func BigBarks(s string) string {
	return dog.WhenGrowUp(s)
}
