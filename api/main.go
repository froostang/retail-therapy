package main

import (
	"embed"
	"fmt"

	"github.com/froostang/retail-therapy/api/cmd"
)

//go:embed build/templates/*
var templateFS embed.FS

func main() {
	fmt.Println("API")

	_, err := templateFS.ReadDir("build/templates")
	if err != nil {
		panic(err)
	}
	cmd.Execute(templateFS)
}
