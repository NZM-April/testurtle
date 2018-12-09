package modules

import(
	"regexp"
	"strings"
	"fmt"
)

func FindTitle(s string) string {
	re := regexp.MustCompile("<title>" + `.*` + "</title>")
	title := re.FindString(s)
	a := strings.Replace(title, "<title>", "", 1)
	trimed_Title := strings.Replace(a, "</title>", "", 1)
	return trimed_Title
}

func (m *Module) TitleModule(){
	if m.Items.Title != "" {
		fmt.Printf("[testurtle] %s : %s\n", m.Items.URL, m.Items.Title)
		title := FindTitle(m.BodyStr)
		b := strings.Contains(title, m.Items.Title)
		m.Judgement(b)
	}
}