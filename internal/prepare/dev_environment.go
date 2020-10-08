package prepare

import (
	"fmt"
	"github.com/dewmal/rakun/internal/utils"
	"github.com/dewmal/rakun/internal/utils/exe_support"
	"log"
	"os/exec"
	"path/filepath"
)

func (env *Environment) BuildDevEnvironment() {
	var err error
	log.Println("Start Build Dev Environment")
	env.EnvPath = env.Config.WorkDir // Set Env Path
	// Create Sample Agent
	log.Println("Number of agents", len(env.Config.Agents))
	for _, agent := range env.Config.Agents {
		log.Println(agent.Name, agent.Code)
		err = utils.PkgFile(fmt.Sprintf("/resources/env_lib_python/%s", agent.Code), filepath.Join(env.EnvPath, agent.Code))
		if err != nil {
			log.Fatalln(err)
		}
	}
	// Copy dev run file to environment
	err = utils.PkgFile("/resources/env_lib_python/run.py", filepath.Join(env.EnvPath, "run.py"))
	if err != nil {
		log.Fatalln(err)
	}
	err = utils.PkgFile("/resources/env_lib_python/requirements.txt", filepath.Join(env.EnvPath, "dev-requirements.txt"))
	if err != nil {
		log.Fatalln(err)
	}
	// Create python virtual environment
	cmd := exec.Command("python", "-m", "venv", filepath.Join(env.EnvPath, "venv"))
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	log.Println(string(stdout))

	exe_support.RunCommand(
		fmt.Sprintf("%s/venv/Scripts/pip.exe", env.EnvPath),
		"install", "-r",
		fmt.Sprintf("%s/dev-requirements.txt", env.EnvPath))

}
