package modules

import(
	"fmt"
	"os/exec"
)

func ShModule(n Notifications){
	if n.Sh != "" {
		script := n.Sh
		err := exec.Command("sh", script).Run()
		if err != nil{
			fmt.Printf("[testurtle] error! %s\n", err)
		} else {
			fmt.Printf("[testurtle] %s is done.", script)
		}
	}
}