package puppy

/*
Adding a tag for v1.0.0
*/

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

func BigBarks(s string) string {
	return dog.WhenGrowUp(s)
}

func From11() {
	fmt.Println("version 1.1.0")

}

func From12() {
	fmt.Println("version 1.2.0")

}
