package runtime

import (
	"context"
	"fmt"
	"log"
	"strings"
)

func (runTime *RunTime) Init() {
	runTime.Context = context.Background()
}

func (runTime *RunTime) ManageAgents() {
	log.Println("Manage Agents ")
	platformChName := fmt.Sprintf("%s_PLATFORM", runTime.Environment.Config.Name)
	platformCtrlChName := fmt.Sprintf("%s_PLATFORM_CTRL", runTime.Environment.Config.Name)
	platformLogChName := fmt.Sprintf("%s_PLATFORM_LOG", runTime.Environment.Config.Name)
	pubsub := runTime.Environment.CommServerClient.Subscribe(runTime.Context, platformChName, platformCtrlChName, platformLogChName)

	ch := pubsub.Channel()
	var initAgents []string
	var isAgentsStarted bool
	for msg := range ch {
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
