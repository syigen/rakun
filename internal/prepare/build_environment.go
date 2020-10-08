package prepare

import (
	"fmt"
	"github.com/dewmal/rakun/internal/utils"
	"log"
	"path/filepath"
)

type Environment struct {
	Config  Config
	EnvPath string
}

func (env *Environment) Build() {
	log.Println("\nStart Build Environment")
	env.EnvPath = fmt.Sprintf("%s/_rakun_env", env.Config.WorkDir) // Set Env Path
	// 1. Create Work Environment
	utils.CreateDir(env.EnvPath, false)
	// 2. Copy required files to environment
	for _, requiredFile := range env.Config.RequiredFiles {
		source := filepath.FromSlash(fmt.Sprintf("%s/%s", env.Config.WorkDir, requiredFile))
		destination := filepath.FromSlash(fmt.Sprintf("%s/%s", env.EnvPath, requiredFile))
		log.Printf("Copying %s to %s\n", source, destination)

		err := utils.Copy(source, destination)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 2. Copy Agent files to environment
	for _, agent := range env.Config.Agents {
		source := filepath.FromSlash(fmt.Sprintf("%s/%s", env.Config.WorkDir, agent.Code))
		destination := filepath.FromSlash(fmt.Sprintf("%s/%s", env.EnvPath, agent.Code))
		log.Printf("Copying %s to %s\n", source, destination)

		err := utils.Copy(source, destination)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 3. Create runtimes
}
