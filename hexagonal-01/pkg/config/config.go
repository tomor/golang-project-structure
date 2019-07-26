package config

const (
	ServiceName = "com.github.tomor.golang-struct-hexagonal-01"
)

// Config contains the service configuration values
type Config struct {
	PostgresURI string `key:"postgres_uri"`
	Kafka KafkaConfig
	Log   log.Config
}

func NewConfig() *Config {
	return new(Config)
}