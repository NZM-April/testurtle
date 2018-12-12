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
			fmt.Println("[testurtle] command done.\n")
		}
	}
}

func ReplaceVariable(cmd string, rn *ResultNums) string {
	okNum := strconv.Itoa(rn.OkNum)
	replacedCmd := strings.Replace(cmd, "$oknum", okNum, 1)
	return replacedCmd
}
