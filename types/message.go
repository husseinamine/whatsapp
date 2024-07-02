package types

// Message object
type Message struct {
	// Required. Messaging service used for the request. Set always to "whatsapp".
	MessagingProduct string `json:"messaging_product"`

	// Optional. Currently, you can only send messages to individuals. Set always to "individual".
	RecipentType string `json:"recipient_type"`

	// Required. WhatsApp ID or phone number for the person you want to send a message to.
	To string `json:"to"`

	/*
		Optional. Only used for Cloud API.
		Used to mention a specific message you are replying to.
		The reply can be any message type currently supported by the Cloud API.
	*/
	Context map[string]string `json:"contextomitempty"`

	// Optional. Default: "text"
	Type MessageType `json:"type,omitempty"`

	// Required for text messages.
	Text Text `json:"text,omitempty"`

	// Required when type is set to audio.
	Audio Media `json:"audio,omitempty"`

	// Required when type is set to contacts.
	Contact Contacts `json:"contact,omitempty"`

	// Required when type is set to document.
	Document Media `json:"document,omitempty"`

	// Required when type is set to image.
	Image Media `json:"image,omitempty"`

	/*
		Required when type is set to interactive.
		An interactive object. This option is used to send List Messages and Reply Buttons.
		The components of each interactive object generally follow a consistent pattern: header, body, footer, and action.
	*/
	Interactive Interactive `json:"interactive,omitempty"`

	// Required when type type is set to location.
	Location Location `json:"location,omitempty"`

	// Required when type type is set to reaction.
	Reaction Reaction `json:"reaction,omitempty"`

	// Required when type is set to sticker.
	Sticker Media `json:"sticker,omitempty"`

	// Required when type is set to video.
	Video Media `json:"video,omitempty"`

	// Required when type is set to template.
	Template Template `json:"template,omitempty"`
}

// MessageType field
type MessageType string

const (
	TextMessage        MessageType = "text"
	TemplateMessage    MessageType = "template"
	DocumentMessage    MessageType = "document"
	ImageMessage       MessageType = "image"
	InteractiveMessage MessageType = "interactive"
	AudioMessage       MessageType = "audio"
	ContactsMessage    MessageType = "contacts"
	LocationMessage    MessageType = "location"
	StickerMessage     MessageType = "sticker"
	VideoMessage       MessageType = "video"
)
