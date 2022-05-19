package servermessages

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type MessageNoData struct {
	Type string `json:"type"`
}

type ErrorPayload struct {
	ID     string      `json:"id,omitempty"`
	Code   string      `json:"code,omitempty"`
	Title  string      `json:"title,omitempty"`
	Detail string      `json:"detail,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

type Challenge struct {
	Payload string `json:"payload"`
}

func CreateClientError(payload ErrorPayload) Message {
	return Message{
		Type: "CLIENT_ERROR",
		Data: payload,
	}
}

func CreateServerError(payload ErrorPayload) Message {
	return Message{
		Type: "SERVER_ERROR",
		Data: payload,
	}
}
