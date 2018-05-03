package main

import "engo.io/engo"

type myScene struct{}

func (*myScene) Type() string { return "myGame" }
func (*myScene) Preload() {
	engo.Files.Load("textures/city.png")
}
func (*myScene) Setup(engo.Updater) {}

func main() {
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  400,
		Height: 400,
	}
	engo.Run(opts, &myScene{})
}
