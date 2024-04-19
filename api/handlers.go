package api

import (
	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go/twiml"
	"log"
	"net/http"
)

const voiceMessage = "Hello, I really believe you should hire Renato. Thanks!"

// postMessage retrieves a message from OpenAI based on the message received from Twilio.
func postMessage(c *gin.Context) {
	body := c.PostForm("Body")
	from := c.PostForm("From")
	to := c.PostForm("To")
	log.Printf("Post request received from %s to %s: %s", from, to, body)

	message, err := getOpenAIMessage(body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("OpenAI message: %s", message)
	m := &twiml.MessagingMessage{
		Body: message,
	}
	response, err := twiml.Messages([]twiml.Element{m})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.Header("Content-Type", "text/xml")
	c.String(http.StatusOK, response)
}

// postVoice responds with a voice message.
func postVoice(c *gin.Context) {
	say := &twiml.VoiceSay{
		Message: voiceMessage,
		Voice:   "Polly.Amy",
	}

	twimlResult, err := twiml.Voice([]twiml.Element{say})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Header("Content-Type", "text/xml")
		c.String(http.StatusOK, twimlResult)
	}
}
