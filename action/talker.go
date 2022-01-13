package action

import "fmt"

type talker interface {
	Talk() string
}

func Perform(t talker) {
	fmt.Println(t.Talk())
}
