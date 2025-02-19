package whatsapp

import (
	"strconv"

	"github.com/husseinamine/whatsapp/types"
)

type TemplateBuilder struct {
	Message *types.Message
	Header  *types.Component
	Body    *types.Component
	Buttons []*types.Component
}

func (mb *TemplateBuilder) WithHeader() *TemplateBuilder {
	mb.Header = &types.Component{
		Type: types.HeaderComponent,
	}

	return mb
}

func (mb *TemplateBuilder) WithHeaderParamater(parameter *types.Parameter) *TemplateBuilder {
	mb.Header.Parameters = append(mb.Header.Parameters, parameter)

	return mb
}

func (mb *TemplateBuilder) WithBody() *TemplateBuilder {
	mb.Body = &types.Component{
		Type: types.BodyComponent,
	}

	return mb
}

func (mb *TemplateBuilder) WithBodyParamater(parameter *types.Parameter) *TemplateBuilder {
	mb.Body.Parameters = append(mb.Body.Parameters, parameter)

	return mb
}

func (mb *TemplateBuilder) WithButton(subtype types.ComponentSubType) *TemplateBuilder {
	buttonCount := len(mb.Buttons)

	if buttonCount >= 3 {
		panic("MORE THAN 3 BUTTONS ADDED TO TEMPLATE")
	}

	mb.Buttons = append(mb.Buttons, &types.Component{
		Type:    types.ButtonComponent,
		SubType: subtype,
		Index:   strconv.Itoa(buttonCount),
	})

	return mb
}

func (mb *TemplateBuilder) WithButtonParameter(parameter *types.ButtonParameter) *TemplateBuilder {
	index := len(mb.Buttons) - 1

	mb.Buttons[index].Parameters = append(mb.Buttons[index].Parameters, parameter)

	return mb
}

func (mb *TemplateBuilder) Save() {
	mb.Message.Template.Components = []*types.Component{
		mb.Header,
		mb.Body,
	}

	mb.Message.Template.Components = append(mb.Message.Template.Components, mb.Buttons...)
}
