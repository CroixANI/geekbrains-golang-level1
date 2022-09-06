package configuration

import (
	"fmt"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

type AppConfigUrl string

func (m *AppConfigUrl) Decode(value string) error {
	parsed, err := url.ParseRequestURI(value)
	if err != nil {
		return err
	}
	*m = AppConfigUrl(parsed.String())
	return nil
}

type AppConfig struct {
	Port        int          `envconfig:"PORT" default:"8080" required:"true"`
	DatabaseUrl AppConfigUrl `envconfig:"DB_URL" required:"true"`
	JaegerUrl   AppConfigUrl `envconfig:"JAEGER_URL" required:"true"`
	SentryUrl   AppConfigUrl `envconfig:"SENTRY_URL" required:"true"`
	KafkaBroker string       `envconfig:"KAFKA_BROKER" required:"true"`
	AppId       string       `envconfig:"SOME_APP_ID" required:"true"`
	AppKey      string       `envconfig:"SOME_APP_KEY" required:"true"`
}

func GetConfiguration() (AppConfig, error) {
	var s AppConfig
	err := envconfig.Process("", &s)
	if err != nil {
		return s, fmt.Errorf("Failed to load configuration: %w", err)
	}

	return s, nil
}
