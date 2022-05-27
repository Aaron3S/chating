package main

import (
	"charting/cmd/server/app"
	"fmt"
)

func main() {
	cmd := app.NewRootCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
