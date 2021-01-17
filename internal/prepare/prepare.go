package prepare

import (
	"github.com/ambelovsky/gosf"
	"github.com/go-redis/redis/v8"
)

type Agent struct {
	Name string
	Code string
}

type AgentList map[string]Agent
type Config struct {
	Name             string
	Version          string
	BuildVersion     int
	Agents           AgentList `yaml:",flow"`
	WorkDir          string
	CommunicationUrl string
	RequiredFiles    []string `yaml:",flow"`
}

type Environment struct {
	Config           *Config
	EnvPath          string
	CommServerClient *redis.Client
}

func (env *Environment) SetupCommServerClient() {
	commUrl := env.Config.CommunicationUrl

	env.CommServerClient = redis.NewClient(&redis.Options{
		Addr:     commUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (env *Environment) SetupDisplayServer() {
	go gosf.Startup(map[string]interface{}{"port": 9999})
}
