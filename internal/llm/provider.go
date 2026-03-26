package llm

import (
	"os"

	"github.com/NunoFrRibeiro/hello-tui/internal/constants"
)

type Selector interface {
	SelectLanguage(userInput string, languages []constants.Language) (int, string, error)
}

type ProviderType string

const (
	ProviderAnthropic ProviderType = "anthropic"
	ProviderOpenAI    ProviderType = "openai"
	ProviderLocal     ProviderType = "local"
	ProviderOllama    ProviderType = "ollama"
)

type Config struct {
	Provider ProviderType
	APIKey   string
	BaseURL  string
	Model    string
}

type providerRule struct {
	Env        string
	MakeConfig func(val string) *Config
}

type ProviderInfo struct {
	Name      string
	Type      ProviderType
	Available bool
	Reason    string
}

var providerRules = []providerRule{
	{
		Env: "ANTHROPIC_API_KEY",
		MakeConfig: func(value string) *Config {
			return &Config{
				Provider: ProviderAnthropic,
				APIKey:   value,
			}
		},
	},
	{
		Env: "OPEN_API_KEY",
		MakeConfig: func(value string) *Config {
			return &Config{
				Provider: ProviderOpenAI,
				APIKey:   value,
			}
		},
	},
	{
		Env: "OLLAMA_HOST",
		MakeConfig: func(value string) *Config {
			return &Config{
				Provider: ProviderOllama,
				BaseURL:  value,
			}
		},
	},
}

func NewSelector(config *Config) Selector {
	if config == nil {
		config = getConfigFromEnv()
	}

	switch config.Provider {
	case ProviderAnthropic:
		return NewAnthropicProvider(config)
	case ProviderOpenAI:
		return NewOpenAIProvider(config)
	case ProviderOllama:
		return NewOllamaProvider(config)
	case ProviderLocal:
		return NewLocalMatcher()
	default:
		return NewLocalMatcher()
	}
}

func getConfigFromEnv() *Config {
	for _, rule := range providerRules {
		if value := os.Getenv(rule.Env); value != "" {
			return rule.MakeConfig(value)
		}
	}
	return &Config{
		Provider: ProviderLocal,
	}
}

func GetProviderName(selector Selector) string {
	switch selector.(type) {
	case *AnthropicProvider:
		return "Anthropic Claude"
	case *OpenAIProvider:
		return "OpenAI GPT"
	case *OllamaProvider:
		return "Ollama (Local)"
	case *LocalMatcher:
		return "Local Pattern Matching"
	default:
		return "Unknown Provider"
	}
}
