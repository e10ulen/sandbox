package main

import "github.com/suganoo/gologger"

func hogeFunc() {
	gologger.Debug("this is debug hogeFunc")
	gologger.Info("call hogeFunc")
}

func main() {
	gologger.SetLogfile("./log")
	defer gologger.CloseFile()

	msg := "hogger"
	//	default debug is muted
	gologger.Debug("this is debug")
	gologger.Info("this is info")
	gologger.Info("msg : " + msg)
	gologger.Warning("this is warning")
	gologger.Error("this is Error")

	gologger.UnmuteDebug()
	hogeFunc()

	gologger.Debug("this is debug xxx")
	gologger.MuteDebug()
	//	this debug message is muted
	gologger.Debug("this is debug yyy")
}
