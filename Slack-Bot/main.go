package main

import (
	"log"
	"github.com/slack-go/slack"
)

func main() {

	OAUTH_TOKEN := "xoxb-3025870033889-2998659128791-ZU77ODBtqDMiuHXtQdLOThmX"
	CHANNEL_ID := "C030FFSQ7JQ"

	api := slack.New(OAUTH_TOKEN)
	attachment := slack.Attachment{
		Pretext: "Pretext",
		Text: "Hello there!",
	}
	
	channelId, timestamp, err := api.PostMessage(
        CHANNEL_ID,
        slack.MsgOptionText("This is the main message", false),
        slack.MsgOptionAttachments(attachment),
        slack.MsgOptionAsUser(true),
    )
 
    if err != nil {
        log.Fatalf("%s\n", err)
    }
 
    log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)
}