package app

type Bridge struct {
	Name      string   `json:"name" mapstructure:"name"`
	Filters   []Filter `json:"filters" mapstructure:"filters"`
	Endpoints []string `json:"endpoints" mapstructure:"endpoints"`
}

func (b *Bridge) Match(msg *Message) bool {
	if len(b.Filters) == 0 {
		return true
	}
	for _, f := range b.Filters {
		if f.Match(msg) {
			return true
		}
	}
	return false
}

type Filter struct {
	To        string `json:"to,omitempty" mapstructure:"to,omitempty"`
	ToRegex   bool   `json:"to_regex,omitempty" mapstructure:"to_regex,omitempty"`
	From      string `json:"from,omitempty" mapstructure:"from,omitempty"`
	FromRegex bool   `json:"from_regex,omitempty" mapstructure:"from_regex,omitempty"`
}

func (f *Filter) Match(msg *Message) bool {
	// TODO: regex
	if f.To != "" {
		if !msg.To[f.To] {
			return false
		}
	}
	if f.From != "" {
		if msg.From != f.From {
			return false
		}
	}
	return true
}
