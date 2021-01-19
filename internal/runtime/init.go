package runtime

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ambelovsky/gosf"
	"log"
	"sort"
	"strings"
)

func (runTime *RunTime) Init() {
	runTime.Context = context.Background()
}

func echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	return gosf.NewSuccessMessage(request.Message.Text)
}

func (runTime *RunTime) ManageAgents() {
	log.Println("Manage Agents ")
	platformChName := fmt.Sprintf("%s_PLATFORM", runTime.Environment.Config.Name)
	platformDisplayChName := fmt.Sprintf("%s_PLATFORM_DISPLAY", runTime.Environment.Config.Name)
	platformCtrlChName := fmt.Sprintf("%s_PLATFORM_CTRL", runTime.Environment.Config.Name)
	platformLogChName := fmt.Sprintf("%s_PLATFORM_LOG", runTime.Environment.Config.Name)
	pubsub := runTime.Environment.CommServerClient.Subscribe(runTime.Context, platformChName, platformCtrlChName, platformLogChName, platformDisplayChName)

	ch := pubsub.Channel()
	var initAgents []string
	var isAgentsStarted bool
	for msg := range ch {
		//log.Println("TEST TEST",msg.Channel, msg.Payload)
		if msg.Channel == platformDisplayChName {
			jsonMap := make(map[string]interface{})

			err := json.Unmarshal([]byte(msg.Payload), &jsonMap)
			if err != nil {
				log.Println(err)
			}
			endPoint := fmt.Sprint(jsonMap["agent"])
			log.Println(jsonMap)
			runTime.Environment.DisplayServer.BroadcastToAll(endPoint, gosf.NewSuccessMessage("Done", jsonMap))
		}

		if msg.Channel == platformCtrlChName {
			if strings.HasPrefix(msg.Payload, "INIT:") {
				agentName := strings.Replace(msg.Payload, "INIT:", "", 1)
				if !contains(initAgents, agentName) {
					initAgents = append(initAgents, agentName)
				}
				if len(runTime.Environment.Config.Agents) == len(initAgents) && !isAgentsStarted {
					for name, _ := range runTime.Environment.Config.Agents {
						log.Println("Agent start ", name)
						runTime.Environment.CommServerClient.Publish(runTime.Context, platformChName, fmt.Sprintf("%s:START", name))
					}
					isAgentsStarted = true
				}
			} else if strings.HasPrefix(msg.Payload, "FAILED:") {
				agentName := strings.Replace(msg.Payload, "INIT:", "", 1)
				log.Println("FAILED AGENT ", agentName)
				runTime.Environment.CommServerClient.Publish(runTime.Context, platformChName, fmt.Sprintf("%s:START", agentName))
			} else if strings.HasPrefix(msg.Payload, "EXIT:") {
				agentName := strings.Replace(msg.Payload, "INIT:", "", 1)
				log.Println("PLATFORM EXIT REQUEST AGENT ", agentName)
				exitCommands := []string{}

				for name, _ := range runTime.Environment.Config.Agents {
					exitCommands = append(exitCommands, fmt.Sprintf("%s:EXIT", name))
				}
				sort.SliceStable(exitCommands, func(i, j int) bool {
					return true
				})
				for _, cmd := range exitCommands {
					//log.Println("Agent exit ", cmd)
					runTime.Environment.CommServerClient.Publish(runTime.Context, platformChName, cmd)
				}
			}
		}
	}

}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
