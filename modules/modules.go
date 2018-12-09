package modules

import(
	"net/http"
)

func ModuleRun(r *http.Response, bodyStr string, i Items, okNum int, ngNum int)(int, int){
	okNum, ngNum = ContainModule(bodyStr, i, okNum, ngNum)
	okNum, ngNum = TitleModule(bodyStr, i, okNum, ngNum)
	okNum, ngNum = StatusModule(r, i, okNum, ngNum)
	return okNum, ngNum
}