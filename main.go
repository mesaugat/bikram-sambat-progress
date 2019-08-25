package main

import (
	"fmt"

	"github.com/tj/go-progress"
)

func main() {
	b := progress.NewInt(100)

	b.Width = 15
	b.StartDelimiter = ""
	b.EndDelimiter = ""

	b.Template(`{{.Bar}} {{.Percent}}%`)

	b.ValueInt(27)

	message := b.String()

	fmt.Println(message)

	Tweet(message)
}
