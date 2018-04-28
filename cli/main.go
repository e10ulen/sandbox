package main

import (
	"log"
	"os"

	"github.com/comail/colog"
	"github.com/spiegel-im-spiegel/astrocalc/internal/mjdnCmd"
	"github.com/spiegel-im-spiegel/gofacade"
)

//	Contextはtodoコマンドのコンテキストを定義
const (
	//	AppNameにはコマンド名を格納
	AppName string = "todo"
	Version string = "0.1.0"
)

func setupFacade(cxt *gofacade.Context) *gofacade.Facade {
	fcd := gofacade.NewFacade(cxt)
	fcd.AddCommand(mjdnCmd.Name, mjdnCmd.Command(cxt, Name))
	return fcd
}

func main() {
	colog.Register()
	cxt := gofacade.NewContext(os.Stdin, os.Stdout, os.Stderr)
	fcd := setupFacade(cxt)
	rtn, err := fcd.Run(Name, Version, os.Args[1:])
	if err != nil {
		cxt.Error(log.Print("e: ", err))
	}
	os.Exit(rtn)

}
