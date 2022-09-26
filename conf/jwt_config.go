package conf

type JWTConfig struct {
	SigningKey string `mapstructure:"signingKey"`
}
