package modules

import(
	"fmt"

	"github.com/NZM-April/testurtle/total"
)

func (m *Module) StatusModule()(int, int){
	if m.Items.Status != 0 {
		var err error
		fmt.Printf("[testurtle] %s : %d\n", m.Items.URL, m.Items.Status)
		if err != nil {
			fmt.Printf("[testurtle] error! %s\n", err)
		}
		b := m.Response.StatusCode == m.Items.Status
		m.OkNum, m.NgNum = total.Judgement(b, m.OkNum, m.NgNum)
	}
	return m.OkNum, m.NgNum
}