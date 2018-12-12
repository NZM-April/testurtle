package modules

import(
	"fmt"
	"os/exec"
	"strings"
	"strconv"

	"github.com/mattn/go-shellwords"
)

func (rn *ResultNums) CommandModule(n Notifications){
	if n.Cmd != "" {
		var out []byte
		replacedCmd := ReplaceVariable(n.Cmd, rn)
		args, err := shellwords.Parse(replacedCmd)
		if err != nil {
			fmt.Println(err)
		}
        out, err = exec.Command(args[0], args[1:]...).Output()
		fmt.Println(string(out))
		if err != nil {
			fmt.Printf("[testurtle] error! %s\n", err)
		} else {
			fmt.Println("[testurtle] command done.")
		}
	}
}

func ReplaceVariable(cmd string, rn *ResultNums) string {
	okNum := strconv.Itoa(rn.OkNum)
	ngNum := strconv.Itoa(rn.NgNum)
	msg := strconv.Itoa(rn.Msg)
	replacedCmd := strings.Replace(cmd, "$oknum", okNum, -1)
	replacedCmd = strings.Replace(cmd, "$ngnum", ngNum, -1)
	replacedCmd = strings.Replace(cmd, "$msg", msg, -1)
	return replacedCmd
}
