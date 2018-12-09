package modules

import(
	"net/http"
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