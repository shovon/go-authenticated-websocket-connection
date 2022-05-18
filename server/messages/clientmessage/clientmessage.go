package clientmessage

import "encoding/json"

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func (m Message) UnmarshalData(v any) error {
	return json.Unmarshal(m.Data, &v)
}

type ChallengeResponse struct {
	Payload   string `json:"payload"`
	Signature string `json:"signature"`
}
