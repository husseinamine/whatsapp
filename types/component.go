package types

// Component object
type Component struct {
	// Required. Describes the component type. For text-based templates, only body is supported.
	Type ComponentType `json:"type"`

	// Required when type is button. Parameter OR ButtonParamater.
	Parameters []*interface{} `json:"parameters,omitempty"`

	// Required when type is button. The type of button to create.
	SubType ComponentSubType `json:"sub_type,omitempty"`

	/*
		Required when type is button. The position index of the button.
		You can have up to 3 buttons using index values of 0 to 2.
	*/
	Index int `json:"index,omitempty"`
}

// ComponentType field
type ComponentType string

const (
	Header ComponentType = "header"
	Body   ComponentType = "body"
	Button ComponentType = "button"
)

// ComponentSubType field
type ComponentSubType string

const (
	QuickReply ComponentSubType = "quick_reply"
	URL        ComponentSubType = "url"
)
