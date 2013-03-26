package cmdexec

import (
	"testing"
    "log"
)

func TestCmdExec(t *testing.T) {
	defer os.Remove("testing.log")
    pgrm := "ps"
    args := "aux"

    c := cmdexec.NewCmdExec(log)
    _ = c.Exec(pgrm+" "+args)
}
