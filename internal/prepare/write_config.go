package prepare

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func (config *Config) InitDevConfig() {
	var err error
	workDir, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}

	config.Name = "Environment Name"
	config.Version = "1.0.0"
	config.BuildVersion = 1
	config.Agents = map[string]Agent{
		"agent1": {
			Name: "Sample Agent One",
			Code: "agent/sample_agent_one.py",
		},
		"agent2": {
			Name: "Sample Agent Two",
			Code: "agent/sample_agent_two.py",
		},
	}
	config.WorkDir = "environment working dir"
	config.RequiredFiles = []string{
		"requirements.txt",
	}

	log.Println("Work Dir ", workDir)
	config.WorkDir = workDir //Set Work Dir
	d, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/RakunConfig", workDir), d, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

}
