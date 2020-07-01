package main

import (
	"fmt"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	t := tokenizer.NewTokenizer()
	morphs, _ := t.Tokenize("すもももももももものうち")
	for i, m := range morphs {
		if m.Id == tokenizer.BOSEOS {
			fmt.Printf("%3d, %v(%v, %v)\n", i, m.Surface, m.Start, m.End)
			continue
		}
		content, _ := m.Content()
		fmt.Printf("%3d, %v(%v, %v)\t%v\n", i, m.Surface, m.Start, m.End, content)
	}
}
