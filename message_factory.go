package whatsapp

import "github.com/husseinamine/whatsapp/types"

type MessageFactory struct {
	Message *types.Message
}

func NewMessageFactory() *MessageFactory {
	mb := &MessageFactory{
		Message: &types.Message{},
	}

	mb.Message.MessagingProduct = "whatsapp"
	mb.Message.RecipentType = "individual"

	return mb
}

func (mb *MessageFactory) WithTo(id string) *MessageFactory {
	mb.Message.To = id

	return mb
}

func (mb *MessageFactory) WithReplyTo(messageId string) *MessageFactory {
	mb.Message.Context["message_id"] = messageId

	return mb
}

func (mb *MessageFactory) WithTemplate(name string, lang string) *TemplateBuilder {
	mb.Message.Type = types.TemplateMessage
	mb.Message.Template = &types.Template{
		Name: name,
	}

	mb.Message.Template.Language = &types.Language{
		Policy: "deterministic",
		Code:   lang,
	}

	return &TemplateBuilder{
		Message: mb.Message,
	}
}
