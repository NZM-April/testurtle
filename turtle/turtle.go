package turtle

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/NZM-April/testurtle/modules"
)

func Start(config string) {
	fmt.Println("[testurtle] Starts!🐢")
	Turtling(config)
}

func Turtling(config string) {
	items, notifications := modules.JsonParse(config)
	Patrol(items)
	Notify(notifications)
}

func Patrol(items []modules.Items) {
	okNum := 0
	ngNum := 0

	for _, i := range items {
		r, err := http.Get(i.URL)
		if err != nil{
			fmt.Println()
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		bodyStr := buf.String()

		okNum, ngNum = modules.ContainModule(bodyStr, i, okNum, ngNum)
		okNum, ngNum = modules.TitleModule(bodyStr, i, okNum, ngNum)
		okNum, ngNum = modules.StatusModule(r, i, okNum, ngNum)
	}
	fmt.Printf("[testurtle] Test completed. OK: \x1b[32m%d\x1b[0m  NG: \x1b[31m%d\x1b[0m\n", okNum, ngNum)
}

func Notify(notifications []modules.Notifications){
	for _, n := range notifications {
		modules.ShModule(n)
	}
}