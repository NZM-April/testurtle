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

var okNum, ngNum int

func Start(config string) {
	fmt.Println("[testurtle] Starts!🐢")
	Turtling(config)
}

func Turtling(config string) {
	items := JsonParse(config)
	Patrol(items)
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
	for _, i := range items {
		r, err := http.Get(i.URL)
		if err != nil{
			fmt.Println()
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		bodyStr := buf.String()

		ContainModule(bodyStr, i)
		TitleModule(bodyStr, i)
		StatusModule(r, i)
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

func ContainModule(bodyStr string, i Items){
	if i.Contain != "" {
		var err error
		fmt.Printf("[testurtle] %s : %s\n", i.URL, i.Contain)
		if err != nil {
			fmt.Printf("[testurtle] error! %s\n", err)
		}
		b := strings.Contains(bodyStr, i.Contain)
		if b == true {
			fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
			okNum++
		} else {
			fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
			ngNum++
		}
	}
}

func TitleModule(bodyStr string, i Items){
	if i.Title != "" {
		var err error
		fmt.Printf("[testurtle] %s : %s\n", i.URL, i.Title)
		if err != nil {
			fmt.Printf("[testurtle] error! %s\n", err)
		}
		title := FindTitle(bodyStr)
		b := strings.Contains(title, i.Title)
		if b == true {
			fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
			okNum++
		} else {
			fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
			ngNum++
		}
	}
}

func StatusModule(r *http.Response, i Items){
	if i.Status != 0 {
		var err error
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