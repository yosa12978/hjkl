package main

import (
	"github.com/yosa12978/hjkl/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
