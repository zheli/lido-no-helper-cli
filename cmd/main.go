package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "myapp",
		Usage: "myapp is a CLI tool",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello, World!")
			return nil
		},
	}
}
