package modules

import(
	"net/http"
	"fmt"
)

type ResultNums struct {
	OkNum int
	NgNum int
	Msg string
}

type Session struct {
	ResultNums *ResultNums
	Response *http.Response
	BodyStr string
	Items Items
}

func (rn *ResultNums) ItemsModuleRun(r *http.Response, bodyStr string, i Items){
	session := Session{rn, r, bodyStr, i}
	session.ContainModule()
	session.TitleModule()
	session.StatusModule()
}

func (s *Session)Judgement(b bool){
	if b == true {
		fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
		s.ResultNums.OkNum++
	} else {
		fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
		s.ResultNums.NgNum++
	}
}

func (rn *ResultNums) NotificationsModuleRun(n Notifications){
	rn.Msg = "Test completed. OK: " + rn.OkNum + "  NG: " + rn.NgNum
	rn.CommandModule(n)
}