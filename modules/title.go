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

func (s *Session) TitleModule(){
	if s.Items.Title != "" {
		fmt.Printf("[testurtle] %s : %s\n", s.Items.URL, s.Items.Title)
		title := FindTitle(s.BodyStr)
		b := strings.Contains(title, s.Items.Title)
		s.Judgement(b)
	}
}