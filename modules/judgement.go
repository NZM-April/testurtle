package modules

import "fmt"

func Judgement(b bool, okNum int, ngNum int)(int, int) {
	if b == true {
		fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
		okNum++
	} else {
		fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
		ngNum++
	}
	return okNum, ngNum
}