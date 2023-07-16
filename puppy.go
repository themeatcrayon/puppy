package puppy

import (
	"fmt"

	"github.com/prinsb/dog"
)

func Bark() string {
	return "Bark!"
}
func Barks() string {
	return "bark bark bark!"
}

func DogSays() {
	fmt.Println(dog.DogModule())
}
