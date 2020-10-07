package prepare

import (
	"fmt"
	"log"
	"os"
)

type Environment struct {
	Config  Config
	EnvPath string
}

func (env *Environment) Build() {
	log.Println("\nStart Build Environment")
	env.EnvPath = fmt.Sprintf("%s/_rakun_env", env.Config.WorkDir) // Set Env Path

	if _, err := os.Stat(env.EnvPath); !os.IsNotExist(err) {
		fmt.Println("\nEnvironment is exists")
	}

	// Create Environment
	if _, err := os.Stat(env.EnvPath); os.IsNotExist(err) {
		err := os.Mkdir(env.EnvPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("\nEnvironment created")
	}

}
