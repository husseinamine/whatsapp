package types

// Language object
type Language struct {
	// Optional. Default: "deterministic"
	Policy string `json:"policy,omitempty"`

	/*
		Required.
		The code of the language or locale to use.
		This field accepts both language (for example, ‘en’) and language_locale (for example, ‘en_US’) formats.
	*/
	Code string `json:"code"`
}
