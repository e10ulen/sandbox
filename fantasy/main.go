package main

import (
	"fmt"
	"log"

	"github.com/comail/colog"
	"github.com/e10ulen/sandbox/lib"
)

func main() {
	colog.Register()
	attribute := []string{"火", "水", "土", "風", "氷", "闇", "聖"}
	spirit := []string{"精霊", "大精霊"}
	create := []string{"矢", "盾", "球"}
	lib.SliceShuffle(attribute)
	lib.SliceShuffle(spirit)
	lib.SliceShuffle(create)
	log.Print("d: 属性:", attribute)
	log.Print("d: 頼む先", spirit)
	log.Print("d: 形:", create)
	fmt.Println(attribute[0] + "の" + spirit[0] + "よ！" + create[0] + "の形を取り")
	if create[0] == "球" || create[0] == "矢" {
		fmt.Println("我が敵になるモノへ飛べ")
	} else if create[0] == "盾" {
		fmt.Println("我が身を守る" + create[0] + "となれ！")
	}
}
