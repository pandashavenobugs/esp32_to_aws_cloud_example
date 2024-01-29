package main

import (
	"fmt"
	"lib/helper"
	"lib/model"
	"lib/service"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

func handler(event model.ThingPayload) {
	helper.EventLogger(event)

	currentTime := time.Now().UTC().Format(time.RFC3339)
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("Error creating session: ", err)
		return
	}
	thingDataDbService := service.NewThingDataDbService(sess)
	thingData := model.ThingData{
		ThingPayload: event,
		CreatedAt:    currentTime,
	}
	err = thingDataDbService.CreateThing(thingData)
	if err != nil {
		fmt.Println("Error creating thing: ", err)
		return
	}

}

func main() {
	lambda.Start(handler)
}
