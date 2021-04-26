package prepare

import (
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/go-redis/redis/v8"
)

type AgentType string

const (
	DynamicAgent AgentType = "dynamic"
	StaticAgent  AgentType = "static"
)

type Agent struct {
	Name string
	Code string
	Type AgentType
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
	DisplayServerUrl string
}

type Environment struct {
	Config           *Config
	EnvPath          string
	CommServerClient *redis.Client
	DisplayServer    *gosocketio.Server
}

const (
	OsWindows = 1
	OsLinux   = 2
)
