package modules

import(
	"net/http"
)

type Session struct {
	OkNum int
	NgNum int
}

type Module struct {
	OkNum int
	NgNum int
	Response *http.Response
	BodyStr string
	Items Items
}

func  (s *Session) ModuleRun(r *http.Response, bodyStr string, i Items){
	module := Module{s.OkNum, s.NgNum, r, bodyStr, i}
	s.OkNum, s.NgNum = module.ContainModule()
	s.OkNum, s.NgNum = module.TitleModule()
	s.OkNum, s.NgNum = StatusModule(r, i, s.OkNum, s.NgNum)
}