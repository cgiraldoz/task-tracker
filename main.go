package main

import "github.com/cgiraldoz/task-tracker/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
