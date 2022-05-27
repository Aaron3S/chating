package main

import (
	"charting/cmd/client/app"
	"fmt"
)

func main() {
	cmd := app.CreateRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
