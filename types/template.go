package types

// Template object
type Template struct {
	// Required.
	Name string `json:"name"`

	// Required.
	Language *Language `json:"language"`

	// Optional.
	Components []*Component `json:"components,omitempty"`
}
