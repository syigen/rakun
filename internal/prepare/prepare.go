package prepare

import (
	"github.com/ambelovsky/gosf"
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
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
	DisplayServer    *gosocketio.Server
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
	log.Panic(http.ListenAndServe("0.0.0.0:9999", serveMux))

}
