package modules

import(
	"fmt"

	"github.com/NZM-April/testurtle/total"
)

func (m *Module) StatusModule(){
	if m.Items.Status != 0 {
		var err error
		fmt.Printf("[testurtle] %s : %d\n", m.Items.URL, m.Items.Status)
		if err != nil {
			fmt.Printf("[testurtle] error! %s\n", err)
		}
		b := m.Response.StatusCode == m.Items.Status
		m.Session.OkNum, m.Session.NgNum = total.Judgement(b, m.Session.OkNum, m.Session.NgNum)
	}
}