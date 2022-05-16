package servermessages

type Message struct {
	Type string      `json:"string"`
	Data interface{} `json:"data"`
}

type ErrorPayload struct {
	ID     string      `json:"id,omitempty"`
	Code   string      `json:"code,omitempty"`
	Title  string      `json:"title,omitempty"`
	Detail string      `json:"detail,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

func CreateClientError(payload ErrorPayload) Message {
	return Message{
		Type: "CLIENT_ERROR",
		Data: payload,
	}
}
