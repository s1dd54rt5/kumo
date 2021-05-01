package main

import (
	"fmt"

	"github.com/spf13/cobra/cobra/cmd"
)

func main() {
	cmd.Execute()
	fmt.Println("Hello")
}
