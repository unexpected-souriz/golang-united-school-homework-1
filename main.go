package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	text := printEmoji()
	fmt.Println(text)
}

func printEmoji() string {
	return emoji.Sprint("Hello :world_map:!")
}
