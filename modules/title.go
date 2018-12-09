package modules

import(
	"regexp"
	"strings"
	"fmt"

	"github.com/NZM-April/testurtle/total"
)

func FindTitle(s string) string {
	re := regexp.MustCompile("<title>" + `.*` + "</title>")
	title := re.FindString(s)
	a := strings.Replace(title, "<title>", "", 1)
	trimed_Title := strings.Replace(a, "</title>", "", 1)
	return trimed_Title
}

func TitleModule(bodyStr string, i Items, okNum int, ngNum int)(int, int){
	if i.Title != "" {
		var err error
		fmt.Printf("[testurtle] %s : %s\n", i.URL, i.Title)
		if err != nil {
			fmt.Printf("[testurtle] error! %s\n", err)
		}
		title := FindTitle(bodyStr)
		b := strings.Contains(title, i.Title)
		okNum, ngNum = total.Judgement(b, okNum, ngNum)
	}
	return okNum, ngNum
}