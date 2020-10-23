package runtime

import (
	"context"
	"github.com/dewmal/rakun/internal/prepare"
)

type RunTime struct {
	Environment *prepare.Environment
	Context     context.Context
}
