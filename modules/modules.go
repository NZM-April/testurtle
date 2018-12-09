package modules

import(
	"net/http"
	"fmt"
)

type ResultNums struct {
	OkNum int
	NgNum int
}

type ModuleArgs struct {
	ResultNums *ResultNums
	Response *http.Response
	BodyStr string
	Items Items
}

func (rn *ResultNums) ModuleRun(r *http.Response, bodyStr string, i Items){
	moduleargs := ModuleArgs{rn, r, bodyStr, i}
	moduleargs.ContainModule()
	moduleargs.TitleModule()
	moduleargs.StatusModule()
}

func (m *ModuleArgs)Judgement(b bool){
	if b == true {
		fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
		m.ResultNums.OkNum++
	} else {
		fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
		m.ResultNums.NgNum++
	}
}