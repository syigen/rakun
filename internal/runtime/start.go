package runtime

import (
	"fmt"
	"github.com/dewmal/rakun/internal/utils/exe_support"
	"log"
	"path/filepath"
)

func (runTime *RunTime) Start() {
	var pythonCommand string
	if OS_LINUX == runTime.getOsType() {
		pythonCommand = ""
	} else {
		pythonCommand = "venv/bin/python"
	}

	exe_support.RunCommand(runTime.buildRuntimeFilePath(pythonCommand), "--version")
	config := runTime.Environment.Config
	agents := config.Agents
	for name, agent := range agents {
		log.Println(fmt.Sprintf("Agent:%s: %s, %s", name, agent.Name, agent.Code))
		commandArgs := []string{
			runTime.buildRuntimeFilePath("run.py"),
			"--stack-name", config.Name,
			"--comm-url", config.CommunicationUrl,
			"--id", name,
			"--name", agent.Name,
			"--source", filepath.ToSlash(agent.Code),
		}
		go exe_support.RunCommand(runTime.buildRuntimeFilePath(pythonCommand), commandArgs...)
	}
}

func (runTime *RunTime) buildRuntimeFilePath(file string) string {
	return filepath.Join(runTime.Environment.EnvPath, file)
}
