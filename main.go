package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

type EchoRequest struct {
	Version string      `json:"version"`
	Session EchoSession `json:"session"`
	Request EchoReqBody `json:"request"`
	Context EchoContext `json:"context"`
}

type EchoSession struct {
	New       bool   `json:"new"`
	SessionID string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

type EchoContext struct {
	System struct {
		Device struct {
			DeviceId string `json:"deviceId,omitempty"`
		} `json:"device,omitempty"`
		Application struct {
			ApplicationID string `json:"applicationId,omitempty"`
		} `json:"application,omitempty"`
	} `json:"System,omitempty"`
}

type EchoReqBody struct {
	Type      string     `json:"type"`
	RequestID string     `json:"requestId"`
	Timestamp string     `json:"timestamp"`
	Intent    EchoIntent `json:"intent,omitempty"`
	Reason    string     `json:"reason,omitempty"`
}

type EchoIntent struct {
	Name  string              `json:"name"`
	Slots map[string]EchoSlot `json:"slots"`
}

type EchoSlot struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Response Types

type EchoResponse struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          EchoRespBody           `json:"response"`
}

type EchoRespBody struct {
	OutputSpeech     *EchoRespPayload `json:"outputSpeech,omitempty"`
	Card             *EchoRespPayload `json:"card,omitempty"`
	Reprompt         *EchoReprompt    `json:"reprompt,omitempty"` // Pointer so it's dropped if empty in JSON response.
	ShouldEndSession bool             `json:"shouldEndSession"`
}

type EchoReprompt struct {
	OutputSpeech EchoRespPayload `json:"outputSpeech,omitempty"`
}

type EchoRespImage struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

type EchoRespPayload struct {
	Type    string        `json:"type,omitempty"`
	Title   string        `json:"title,omitempty"`
	Text    string        `json:"text,omitempty"`
	SSML    string        `json:"ssml,omitempty"`
	Content string        `json:"content,omitempty"`
	Image   EchoRespImage `json:"image,omitempty"`
}

func HandleRequest(ctx context.Context, req EchoRequest) (EchoResponse, error) {
	logrus.Infof("REQ: %+v", req)

	intentName := req.Request.Intent.Name

	logrus.Infof("Intent: %s", intentName)

	res := EchoResponse{
		Version: "1.0",
		Response: EchoRespBody{
			OutputSpeech: &EchoRespPayload{
				Type: "PlainText",
				Text: "This is a test",
			},
			ShouldEndSession: true,
		},
	}

	logrus.Infof("Res: %+v", res)

	return res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
