package prepare

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
	Config  *Config
	EnvPath string
}
