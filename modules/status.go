package modules

import(
	"fmt"
)

func (s *Session) StatusModule(){
	if s.Items.Status != 0 {
		fmt.Printf("[testurtle] %s : %d\n", s.Items.URL, s.Items.Status)
		b := s.Response.StatusCode == s.Items.Status
		s.Judgement(b)
	}
}