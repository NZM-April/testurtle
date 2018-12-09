package modules

import(
	"fmt"
	"strings"

	"github.com/NZM-April/testurtle/total"
)

func (m *Module) ContainModule()(int, int){
	if m.Items.Contain != "" {
		fmt.Printf("[testurtle] %s : %s\n", m.Items.URL, m.Items.Contain)
		b := strings.Contains(m.BodyStr, m.Items.Contain)
		m.OkNum, m.NgNum = total.Judgement(b, m.OkNum, m.NgNum)
	}
	return m.OkNum, m.NgNum
}