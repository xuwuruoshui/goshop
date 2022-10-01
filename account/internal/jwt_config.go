package internal

type JWT struct {
	SigningKey string `mapstructure:"signingKey" yaml:"signingKey"`
}
