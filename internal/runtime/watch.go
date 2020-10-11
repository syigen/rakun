package runtime

import (
	"context"
	"log"
)

var ctx = context.Background()

func (runTime *RunTime) Watch() {
	pubsub := runTime.Environment.CommServerClient.Subscribe(ctx, "mychannel1")

	ch := pubsub.Channel()

	for msg := range ch {
		log.Println(msg.Channel, msg.Payload)
	}
}
