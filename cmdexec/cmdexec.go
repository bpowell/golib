package cmdexec

import(
    "os/exec"
    "bytes"
    "log"
    "strings"
)

type CmdExec struct {
    log *log.Logger
}

func NewCmdExec(log *log.Logger) *CmdExec {
    return &CmdExec{log: log}
}

func (c *CmdExec) Exec(command string) bool {
    var out bytes.Buffer
    s := strings.Split(command, " ")
    pgrm := s[0]
    args := s[1:]

    cmd := exec.Command(pgrm, args...)
    cmd.Stdout = &out

    err := cmd.Run()
    if err != nil {
        c.log.Fatalln("Cannot run command" + command)
    }

    c.log.Println(out.String())

    return (cmd.ProcessState).Success()
}
