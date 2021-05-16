package main

import (
	"context"

	"github.com/eihigh/goban"
)

var (
	grid = goban.NewGrid(
		"		1fr		1fr		1fr",
		"1fr	header	header	header",
		"3fr side		content	",
		"1fr	footer	footer	footer",
	)
)

func main() {
	goban.Main(app, view)
}

func app(_ context.Context, es goban.Events) error {
	goban.Show()
	es.ReadKey()
	return nil
}

func view() {
	b := goban.Screen().Enclose("")
	header := b.GridItem(grid, "header").DrawSides("", 0, 0, 0, 1)
	header.Prints("Header")
}
