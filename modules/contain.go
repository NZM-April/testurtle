package modules

import(
	"fmt"
	"strings"
)

func ContainModule(bodyStr string, i Items, okNum int, ngNum int)(int, int){
	if i.Contain != "" {
		fmt.Printf("[testurtle] %s : %s\n", i.URL, i.Contain)
		b := strings.Contains(bodyStr, i.Contain)
		okNum, ngNum = Judgement(b, okNum, ngNum)
	}
	return okNum, ngNum
}