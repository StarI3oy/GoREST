package gorest

import (
	"fmt"
)

type user struct {
	ID    string `json:"id"`
	name  string `json:"name"`
	title string `json:"title"`
}

func main() {
	fmt.Println("hello world")

}
