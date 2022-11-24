package main

import "github.com/alecthomas/kong"

var CLI struct {
	Rename struct {
		Name struct {
			Name string `arg` // <-- NOTE: identical name to enclosing struct field.
			To   struct {
				Name struct {
					Name string `arg`
				} `arg`
			} `cmd`
		} `arg`
	} `cmd`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "rm <path>":
	case "ls":
	default:
		panic(ctx.Command())
	}
}
