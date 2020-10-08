package prepare

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func (config *Config) InitRunConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	log.Println("Work Dir ", workDir)
	config.WorkDir = workDir //Set Work Dir

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/RakunConfig", workDir))
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(content), config)
	if err != nil {
		log.Println("error: %v", err)
	}
	log.Println(fmt.Sprintf("Name: %s", config.Name))
	log.Println(fmt.Sprintf("Version: %s", config.Version))
	log.Println(fmt.Sprintf("Build Version: %d", config.BuildVersion))
	log.Println("Agents")
	for agentKey, agent := range config.Agents {
		log.Println(fmt.Sprintf("id: %s", agentKey))
		log.Println(fmt.Sprintf("\tname: %s", agent.Name))
		log.Println(fmt.Sprintf("\tcode: %s", agent.Code))
	}
	// Required Files
	log.Println("Required Files")
	for _, file := range config.RequiredFiles {
		log.Println(file)
	}

}

func (_ *Config) GenConfigSample() {
	sam := Config{
		Name:         "Environment Name",
		Version:      "",
		BuildVersion: 0,
		Agents: map[string]Agent{
			"agent1": {
				Name: "A",
				Code: "A.py",
			},
			"agent2": {
				Name: "B",
				Code: "B.py",
			},
		},
		WorkDir: "environment working dir",
		RequiredFiles: []string{
			"a.txt",
			"main.py",
		},
	}
	d, err := yaml.Marshal(&sam)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- m dump:\n%s\n\n", string(d))
}
