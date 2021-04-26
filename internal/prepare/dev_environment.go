package prepare

import (
	"fmt"
	"github.com/ambelovsky/gosf"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"github.com/dewmal/rakun/internal/utils"
	"github.com/dewmal/rakun/internal/utils/exe_support"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
)

func (env *Environment) BuildDevEnvironment() {
	var err error
	log.Println("Start Build Dev Environment")
	var pythonCommand string
	if OsLinux == env.GetOsType() {
		pythonCommand = "venv/bin/pip"
	} else {
		pythonCommand = "venv/Scripts/pip.exe"
	}

	env.EnvPath = env.Config.WorkDir // Set Env Path

	err = utils.PkgFile("/resources/env_lib_python/agent/__init__.py", filepath.Join(env.EnvPath, "agent/__init__.py"))
	if err != nil {
		log.Fatalln(err)
	}
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
		fmt.Sprintf("%s/%s", env.EnvPath, pythonCommand),
		"install", "-r",
		fmt.Sprintf("%s/dev-requirements.txt", env.EnvPath))

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
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	env.DisplayServer = server

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel, args interface{}) {
		//client id is unique
		log.Println("New client connected, client id is ", c.Id())

		//you can join clients to rooms
		//c.Join("rakun")
		c.Emit("test", gosf.NewSuccessMessage("Test"))

		log.Println(c.IsAlive())

	})
	server.On("send", func(c *gosocketio.Channel, msg gosf.Message) string {
		//send event to all in room
		c.BroadcastTo("chat", "message", msg)
		return "OK"
	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)

	log.Println("Display server started ", env.Config.DisplayServerUrl)
	error := http.ListenAndServe(env.Config.DisplayServerUrl, serveMux)
	log.Panic(error)

}

func (env *Environment) GetOsType() int {
	os := runtime.GOOS
	switch os {
	case "windows":
		return OsWindows
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		return OsLinux
	default:
		fmt.Printf("%s.\n", os)
	}
	return 0
}
