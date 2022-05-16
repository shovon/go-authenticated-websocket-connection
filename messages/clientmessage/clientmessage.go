package clientmessage

import "encoding/json"

type Message struct {
	Type string          `json:"string"`
	Data json.RawMessage `json:"data"`
}

type ChallengeResponse struct {
	Payload   string `json:"payload"`
	Signature string `json:"string"`
}
