package modules

import(
	"fmt"
	"strings"
)

func (m *ModuleArgs) ContainModule(){
	if m.Items.Contain != "" {
		fmt.Printf("[testurtle] %s : %s\n", m.Items.URL, m.Items.Contain)
		b := strings.Contains(m.BodyStr, m.Items.Contain)
		m.Judgement(b)
	}
}