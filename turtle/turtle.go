package turtle

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/NZM-April/testurtle/modules"
)

var resultNums modules.ResultNums

func Start(config string) {
	fmt.Println("[testurtle] Starts!🐢")
	Turtling(config)
}

func Turtling(config string) {
	items, notifications := modules.JsonParse(config)
	ItemsRun(items)
	NotificationsRun(notifications)
}

func ItemsRun(items []modules.Items) {
	for _, i := range items {
		r, err := http.Get(i.URL)
		if err != nil{
			fmt.Println(err)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		bodyStr := buf.String()

		resultNums.ItemsModuleRun(r, bodyStr, i)
	}
	fmt.Printf("[testurtle] Test completed. OK: \x1b[32m%d\x1b[0m  NG: \x1b[31m%d\x1b[0m\n", resultNums.OkNum, resultNums.NgNum)
}

func NotificationsRun(notifications []modules.Notifications){
	for _, n := range notifications {
		resultNums.NotificationsModuleRun(n)
	}
}
