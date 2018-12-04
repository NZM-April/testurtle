package turtle

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"bytes"
)

type Check struct {
    URL string `json:"URL"`
    Target string `json:"target"`
}

func Start(config string) {
	fmt.Println("[testurtle] Starts!🐢")
	Turtling(config)
}

func Turtling(configFile string){
	if config != ""{
		configFile = c
	} else {
		configFile = "turtleconfig.json"
	}
	bytes, err := ioutil.ReadFile(config)
    if err != nil {
        log.Fatal(err)
    }
    var checks []Check
    if err := json.Unmarshal(bytes, &checks); err != nil {
        log.Fatal(err)
    }
	Patrol(checks)
}

func Patrol(checks []Check){
	for _, c := range checks {
		fmt.Printf("[testurtle] %s : %s\n", c.URL, c.Target)
		r, err := http.Get(c.URL)
		if err != nil{
			log.Fatal(err)
		}
		buf := new(bytes.Buffer)
    	buf.ReadFrom(r.Body)
    	newStr := buf.String()
		b := strings.Contains(newStr, c.Target)
		if b == true{
			fmt.Println("[testurtle] => OK")
		}
    }
}
