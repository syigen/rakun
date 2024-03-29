package prepare

import (
	"fmt"
	"github.com/dewmal/rakun/internal/utils"
	"github.com/dewmal/rakun/internal/utils/exe_support"
	"log"
	"os/exec"
	"path/filepath"
)

func (env *Environment) BuildRunEnvironment() {
	var err error
	log.Println("\nStart Build Environment")
	var pythonCommand string
	if OsLinux == env.GetOsType() {
		pythonCommand = "venv/bin/pip"
	} else {
		pythonCommand = "venv/Scripts/pip.exe"
	}

	env.EnvPath = fmt.Sprintf("%s/_rakun_env", env.Config.WorkDir) // Set Env Path
	// 1. Create Work Environment
	utils.CreateDir(env.EnvPath, false)

	// 2. Copy required files to environment
	for _, requiredFile := range env.Config.RequiredFiles {
		source := filepath.FromSlash(fmt.Sprintf("%s/%s", env.Config.WorkDir, requiredFile))
		destination := filepath.FromSlash(fmt.Sprintf("%s/%s", env.EnvPath, requiredFile))
		log.Printf("Copying %s to %s\n", source, destination)

		err = utils.Copy(source, destination)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 2. Copy Agent files to environment
	log.Println("Number of agents ", len(env.Config.Agents))
	for name, agent := range env.Config.Agents {
		log.Println("Agent Name", name)
		source := filepath.FromSlash(fmt.Sprintf("%s/%s", env.Config.WorkDir, agent.Code))
		destination := filepath.FromSlash(fmt.Sprintf("%s/%s", env.EnvPath, agent.Code))
		log.Printf("Copying %s to %s\n", source, destination)

		err = utils.Copy(source, destination)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 3. Create runtime
	// 3.1 Add runtime files
	err = utils.PkgFile("/resources/env_lib_python/run.py", fmt.Sprintf("%s/run.py", env.EnvPath))
	if err != nil {
		log.Fatalln(err)
	}
	err = utils.PkgFile("/resources/env_lib_python/requirements.txt", fmt.Sprintf("%s/requirements.txt", env.EnvPath))
	if err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("python", "-m", "venv", fmt.Sprintf("%s/venv", env.EnvPath))
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	log.Println(string(stdout))

	exe_support.RunCommand(
		fmt.Sprintf("%s/%s", env.EnvPath, pythonCommand),
		"install", "-r",
		fmt.Sprintf("%s/requirements.txt", env.EnvPath))

}
