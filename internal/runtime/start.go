package runtime

import (
	"github.com/dewmal/rakun/internal/utils/exe_support"
	"path/filepath"
)

func (runTime *RunTime) Start() {
	exe_support.RunCommand(runTime.buildRuntimeFilePath("venv/Scripts/python.exe"), "--version")
	exe_support.RunCommand(runTime.buildRuntimeFilePath("venv/Scripts/python.exe"), runTime.buildRuntimeFilePath("run.py"))
}

func (runTime *RunTime) buildRuntimeFilePath(file string) string {
	return filepath.Join(runTime.Environment.EnvPath, file)
}
