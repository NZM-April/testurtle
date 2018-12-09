package modules

import(
	"fmt"
	"strings"
)

func (s *Session) ContainModule(){
	if s.Items.Contain != "" {
		fmt.Printf("[testurtle] %s : %s\n", s.Items.URL, s.Items.Contain)
		b := strings.Contains(s.BodyStr, s.Items.Contain)
		s.Judgement(b)
	}
}