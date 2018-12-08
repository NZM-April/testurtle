package modules

import(
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
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