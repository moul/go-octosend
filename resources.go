package octosend

type SpoolerMessage struct {
	Attachments []string          `json:"attachments"`
	Headers     map[string]string `json:"headers"`
	Parts       []string          `json:"parts"`
	Recipient   string            `json:"recipient"`
	Sender      string            `json:"sender"`
	Subject     string            `json:"subject"`
	Unsubscribe string            `json:"unsubscribe"`
	Variables   map[string]string `json:"variables"`
}

type Domain struct {
	Name string `json:"name"`
}
