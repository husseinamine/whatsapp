package types

// Parameter object.
type Parameter struct {
	// Required.
	Type ParameterType `json:"type"`

	// Required when Type is "text".
	// For the header component, the character limit is 60 characters.
	// For the body component, the character limit is 1024.
	// The exception to these character limits applies to template messages in the following conditions:
	// - When sending a template message with a body component only, the character limit for the text parameter and the full template text is 32768 characters.
	// - When sending a template message with body and other components, the character limit for the text parameter and the full template text is 1024 characters.
	Text string `json:"text,omitempty"`

	// Required when Type is "currency".
	Currency *Currency `json:"currency,omitempty"`

	// Required when Type is "date_time".
	DateTime *DateTime `json:"date_time,omitempty"`

	// Required when Type is "image".
	Image *Media `json:"image,omitempty"`

	// Required when Type is "document".
	// Only PDF documents are supported for media-based message templates.
	Document *Media `json:"document,omitempty"`
}

type ParameterType string

const (
	TextParameter     ParameterType = "text"
	CurrencyParameter ParameterType = "currency"
	DateTimeParameter ParameterType = "date_time"
	ImageParameter    ParameterType = "image"
	DocumentParameter ParameterType = "document"
)
