package prepare

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Agent struct {
	Name string
	Code string
}

type AgentList map[string]Agent

type Config struct {
	Name         string
	Version      string
	BuildVersion int
	Agents       AgentList `yaml:",flow"`
}

func (config *Config) Init() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Work Dir ", workDir)
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/RakunConfig", workDir))
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(content), config)
	if err != nil {
		log.Println("error: %v", err)
	}
	fmt.Println(fmt.Sprintf("Name: %s\nVersion:%s\nBuild Version:%d\nAgents", config.Name, config.Version, config.BuildVersion))
	for agentKey, agent := range config.Agents {
		fmt.Println(fmt.Sprintf("\tid: %s\n\t\tname: %s\n\t\tcode: %s", agentKey, agent.Name, agent.Code))
	}
}
