package main

import (
	"context"

	"github.com/eihigh/goban"
)

func main() {
	goban.Main(app)
}

func app(_ context.Context, es goban.Events) error {
	v := func() {
		b := goban.Screen()
		b.Puts("Press any key to open popup")
		b.Puts("Ctrl+C to exit.")
	}
	goban.PushViewFunc(v)

	for {
		goban.Show()
		es.ReadKey()
		popup(es)
	}
}

func popup(es goban.Events) {
	//  ポップアップ処理
	v := func() {
		b := goban.NewBox(0, 0, 40, 5).Enclose("popup window")
		b.Prints("Press any key to close Popup")
	}

	goban.PushViewFunc(v)
	defer goban.PopView()

	goban.Show()
	es.ReadKey()
}

/*
func app(_ context.Context, es goban.Events) error {
  goban.Show()
  es.ReadKey()
  return nil
}

func view() {
  goban.Screen().Enclose("hello").Prints("Hello World\nPress any key to exit.")
}
*/
