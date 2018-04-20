package main

import (
	"fmt"
)
func check(ping bool, res string) (bool ,string){
	res = "check"
	ping = false
	if ping != true {
		res = res+"true"
		return ping, res
	}else{
		res = res+"false"
		return ping, res
	}
}

func main(){
	fmt.Println("Test Method")
	var get string
	var pong bool
	fmt.Println(check(pong, get))
}
