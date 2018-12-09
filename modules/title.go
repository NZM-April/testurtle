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

func (m *Module) TitleModule()(int, int){
	if m.Items.Title != "" {
		var err error
		fmt.Printf("[testurtle] %s : %s\n", m.Items.URL, m.Items.Title)
		if err != nil {
			fmt.Printf("[testurtle] error! %s\n", err)
		}
		title := FindTitle(m.BodyStr)
		b := strings.Contains(title, m.Items.Title)
		m.OkNum, m.NgNum = total.Judgement(b, m.OkNum, m.NgNum)
	}
	return m.OkNum, m.NgNum
}