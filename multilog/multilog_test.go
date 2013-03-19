package multilog

import (
	"testing"
)

func TestNewMultiLog(t *testing.T) {
	defer os.Remove("testing.log")

	a, _ = multilog.NewMultiLog("Logging", 6, "testing.log")

	a.Println("Hello from TestNewMultiLog")
}
