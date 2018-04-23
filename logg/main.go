package main

import (
	"log"
	"os"
)

func main() {
	debug := log.New(os.Stdout, "[Debug]", log.LstdFlags|log.Lshortfile)
	warn := log.New(os.Stdout, "[WARN]", log.LstdFlags)
	debug.Print("debug test")
	warn.Print("warn test")
}
