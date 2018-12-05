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

type Check struct {
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
	var configFile string
	var checks []Check

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

	if err := json.Unmarshal(bytes, &checks); err != nil {
		fmt.Printf("[testurtle] error! %s\n", err)
		os.Exit(1)
	}
	
	Patrol(checks)
}

func Patrol(checks []Check) {
	okNum := 0
	ngNum := 0
	for _, c := range checks {
		r, err := http.Get(c.URL)
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		newStr := buf.String()

		if c.Contain != "" {
			fmt.Printf("[testurtle] %s : %s\n", c.URL, c.Contain)
			if err != nil {
				fmt.Printf("[testurtle] error! %s\n", err)
			}
			b := strings.Contains(newStr, c.Contain)
			if b == true {
				fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
				okNum++
			} else {
				fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
				ngNum++
			}
		}

		if c.Title != "" {
			fmt.Printf("[testurtle] %s : %s\n", c.URL, c.Title)
			if err != nil {
				fmt.Printf("[testurtle] error! %s\n", err)
			}
			title := FindTitle(newStr)
			b := strings.Contains(title, c.Title)
			if b == true {
				fmt.Printf("%s \x1b[32m%s\x1b[0m\n", "[testurtle] =>", "ok")
				okNum++
			} else {
				fmt.Printf("%s \x1b[31m%s\x1b[0m\n", "[testurtle] =>", "ng")
				ngNum++
			}
		}

		if c.Status != 0 {
			fmt.Printf("[testurtle] %s : %d\n", c.URL, c.Status)
			if err != nil {
				fmt.Printf("[testurtle] error! %s\n", err)
			}
			if r.StatusCode == c.Status {
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
