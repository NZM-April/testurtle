package modules

import(
	"fmt"
)

func (m *ModuleArgs) StatusModule(){
	if m.Items.Status != 0 {
		fmt.Printf("[testurtle] %s : %d\n", m.Items.URL, m.Items.Status)
		b := m.Response.StatusCode == m.Items.Status
		m.Judgement(b)
	}
}