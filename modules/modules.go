package modules

import(
	"net/http"
	"fmt"
)

type Session struct {
	OkNum int
	NgNum int
}

type Module struct {
	Session *Session
	Response *http.Response
	BodyStr string
	Items Items
}

func (s *Session) ModuleRun(r *http.Response, bodyStr string, i Items){
	module := Module{s, r, bodyStr, i}
	module.ContainModule()
	module.TitleModule()
	module.StatusModule()
}

func (m *Module)Judgement(b bool){
	if b == true {
		fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
		m.Session.OkNum++
	} else {
		fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
		m.Session.NgNum++
	}
}