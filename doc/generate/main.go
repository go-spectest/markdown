//go:build linux || darwin

// Package main is generating markdown.
package main

import (
	"os"

	md "github.com/go-spectest/markdown"
)

//go:generate go run main.go

func main() {
	f, err := os.Create("generated.md")
	if err != nil {
		panic(err)
	}

	if err := md.NewMarkdown(f).
		H1("go generate example").
		PlainText("This markdown is generated by `go generate`").
		Build(); err != nil {
		panic(err)
	}
}
