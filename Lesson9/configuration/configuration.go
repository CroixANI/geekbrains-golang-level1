package configuration

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type SourceConfigType int

const (
	EnvConfig SourceConfigType = iota
	JasonConfig
	YamlConfig
)

type AppConfigUrl url.URL

func (m *AppConfigUrl) Decode(value string) error {
	val, err := url.Parse(value)
	if err != nil {
		return err
	}
	*m = AppConfigUrl(*val)
	return nil
}

func (m *AppConfigUrl) UnmarshalJSON(value []byte) error {
	val, err := url.Parse(strings.Replace(string(value), "\"", "", -1))
	if err != nil {
		return err
	}
	*m = AppConfigUrl(*val)
	return nil
}

func (m *AppConfigUrl) UnmarshalYAML(value *yaml.Node) error {
	val, err := url.Parse(strings.Replace(value.Value, "\"", "", -1))
	if err != nil {
		return err
	}
	*m = AppConfigUrl(*val)
	return nil
}

type AppConfig struct {
	Port        int          `envconfig:"PORT" default:"8080" required:"true" json:"port" yaml:"port"`
	DatabaseUrl AppConfigUrl `envconfig:"DB_URL" required:"true" json:"dbUrl" yaml:"dbUrl"`
	JaegerUrl   AppConfigUrl `envconfig:"JAEGER_URL" required:"true" json:"jaegerUrl" yaml:"jaegerUrl"`
	SentryUrl   AppConfigUrl `envconfig:"SENTRY_URL" required:"true" json:"sentryUrl" yaml:"sentryUrl"`
	KafkaBroker string       `envconfig:"KAFKA_BROKER" required:"true" json:"kafkaBroker" yaml:"kafkaBroker"`
	AppId       string       `envconfig:"SOME_APP_ID" required:"true" json:"appId" yaml:"appId"`
	AppKey      string       `envconfig:"SOME_APP_KEY" required:"true" json:"appKey" yaml:"appKey"`
}

func GetConfiguration(sourceType SourceConfigType) (*AppConfig, error) {
	switch sourceType {
	case EnvConfig:
		return getEnvConfig()
	case JasonConfig:
		return getJsonConfig()
	case YamlConfig:
		return getYamlConfig()
	default:
		return nil, fmt.Errorf("unsupported config source type: %v", sourceType)
	}
}

func getEnvConfig() (*AppConfig, error) {
	var s AppConfig
	err := envconfig.Process("", &s)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return &s, nil
}

func getJsonConfig() (*AppConfig, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config AppConfig
	if err = json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func getYamlConfig() (*AppConfig, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config AppConfig
	if err = yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return &config, nil
}
