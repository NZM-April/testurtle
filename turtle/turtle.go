package turtle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"regexp"
)

type Modules struct {
	Items []Items `json:"items"`
}

type Items struct {
	URL    string `json:"URL"`
	Contain string `json:"contain"`
	Title string `json:"title"`
	Status int `json:"status"`
}

func Start(config string) {
	fmt.Println("[testurtle] Starts!🐢")
	Turtling(config)
}

func Turtling(config string) {
	fmt.Println(JsonParse(config))
	//Patrol(items)
}

func JsonParse(config string) []Items{
	var configFile string
	var modules Modules

	if config != "" {
		configFile = config
	} else {
		configFile = "turtleconfig.json"
	}
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("[testurtle] error! %s\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(bytes, &modules); err != nil {
		fmt.Printf("[testurtle] error! %s\n", err)
		os.Exit(1)
	}
	return modules.Items
}

func Patrol(items []Items) {
	okNum := 0
	ngNum := 0
	for _, i := range items {
		r, err := http.Get(i.URL)
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		newStr := buf.String()

		if i.Contain != "" {
			fmt.Printf("[testurtle] %s : %s\n", i.URL, i.Contain)
			if err != nil {
				fmt.Printf("[testurtle] error! %s\n", err)
			}
			b := strings.Contains(newStr, i.Contain)
			if b == true {
				fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
				okNum++
			} else {
				fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
				ngNum++
			}
		}

		if i.Title != "" {
			fmt.Printf("[testurtle] %s : %s\n", i.URL, i.Title)
			if err != nil {
				fmt.Printf("[testurtle] error! %s\n", err)
			}
			title := FindTitle(newStr)
			b := strings.Contains(title, i.Title)
			if b == true {
				fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
				okNum++
			} else {
				fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
				ngNum++
			}
		}

		if i.Status != 0 {
			fmt.Printf("[testurtle] %s : %d\n", i.URL, i.Status)
			if err != nil {
				fmt.Printf("[testurtle] error! %s\n", err)
			}
			if r.StatusCode == i.Status {
				fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
				okNum++
			} else {
				fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
				ngNum++
			}
		}
	}
	fmt.Printf("[testurtle] Test completed. OK: \x1b[32m%d\x1b[0m  NG: \x1b[31m%d\x1b[0m\n", okNum, ngNum)
}

func FindTitle(s string) string{
	re := regexp.MustCompile("<title>" + `.*` + "</title>")
	title := re.FindString(s)
	a := strings.Replace(title, "<title>", "", 1)
	trimed_Title := strings.Replace(a, "</title>", "", 1)
	return trimed_Title
}
