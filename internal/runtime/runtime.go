package runtime

import (
	"context"
	"fmt"
	"github.com/dewmal/rakun/internal/prepare"
	"runtime"
)

const (
	OS_WINDOWS = 1
	OS_LINUX   = 2
)

type RunTime struct {
	Environment *prepare.Environment
	Context     context.Context
}

func (runTime RunTime) getOsType() int {
	os := runtime.GOOS
	switch os {
	case "windows":
		return OS_WINDOWS
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		return OS_LINUX
	default:
		fmt.Printf("%s.\n", os)
	}
	return 0
}
