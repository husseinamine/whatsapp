package types

// ButtonParameter object.
type ButtonParameter struct {
	// Required.
	Type ButtonParameterType `json:"type"`

	// Required when Type is "payload".
	Payload string `json:"payload,omitempty"`

	// Required when Type is "date_time".
	Text string `json:"text,omitempty"`
}

type ButtonParameterType string

const (
	TextButtonParameter    ButtonParameterType = "text"
	PayloadButtonParameter ButtonParameterType = "payload"
)
