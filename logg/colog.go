package main

import (
	"fmt"
	"log"
	"os"

	"github.com/comail/colog"
)

//	logの省略
//	t: trace
//	d: debug
//	i: info
//	w: warning
//	e: error
//	a: alert
func logMap() {
	log.Print("t: trace")
	log.Print("trc: trace")
	log.Print("trace: trace")
	log.Print("d: debug")
	log.Print("dbg: debug")
	log.Print("debug: debug")
	log.Print("i: info")
	log.Print("inf: info")
	log.Print("info: info")
	log.Print("w: warning")
	log.Print("wrn: warning")
	log.Print("warn: warning")
	log.Print("warning: warning")
	log.Print("e: error")
	log.Print("err: error")
	log.Print("error: error")
	log.Print("a: alert")
	log.Print("alr: alert")
	log.Print("alert: alert")
	log.Print("panic: alert")
}

func logLevel() {
	fmt.Println("set log level")
	cl := colog.NewCoLog(os.Stdout, "worker", log.LstdFlags)
	cl.SetMinLevel(colog.LInfo)
	cl.SetDefaultLevel(colog.LWarning)
	cl.FixedValue("worker_id", 42)

	logger := cl.NewLogger()
	logger.Print("this gets warning level")
	logger.Print("debug: this won't be displayed")
}

func logTest() {
	colog.SetMinLevel(colog.LInfo)
	colog.SetDefaultLevel(colog.LWarning)
}

func main() {
	colog.Register()

	logMap()
	logLevel()
	logTest()
	fmt.Println("set MinLevel Info")
	logMap()
}
