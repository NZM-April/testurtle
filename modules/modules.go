package modules

import(
	"net/http"
)

type Session struct {
	OkNum int
	NgNum int
}

func  (s *Session) ModuleRun(r *http.Response, bodyStr string, i Items){
	s.OkNum, s.NgNum = ContainModule(bodyStr, i, s.OkNum, s.NgNum)
	s.OkNum, s.NgNum = TitleModule(bodyStr, i, s.OkNum, s.NgNum)
	s.OkNum, s.NgNum = StatusModule(r, i, s.OkNum, s.NgNum)
}