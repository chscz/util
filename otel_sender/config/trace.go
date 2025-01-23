package config

type Trace struct {
	Enable            bool              `env:"ENABLED"`
	ContentType       ContentType       `env:"CONTENT_TYPE"`
	ProtocolType      ProtocolType      `env:"PROTOCOL_TYPE"`
	TerminationConfig TerminationConfig `envPrefix:"TERMINATION_CONFIG_"`
}
