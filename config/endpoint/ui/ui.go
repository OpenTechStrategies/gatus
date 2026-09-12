package ui

import "errors"

// Config is the UI configuration for endpoint.Endpoint
type Config struct {
	// RecentChecksMaximumRows is the maximum number of rows displayed for recent checks.
	RecentChecksMaximumRows int `yaml:"recent-checks-maximum-rows,omitempty"`
	// RecentChecksResultsPerRow is the number of results displayed in each Recent Checks row.
	RecentChecksResultsPerRow int `yaml:"recent-checks-results-per-row,omitempty"`
	// RecentChecksResultHeight is the CSS height applied to each Recent Checks result.
	RecentChecksResultHeight string `yaml:"recent-checks-result-height,omitempty"`
	// HideConditions whether to hide the condition results on the UI
	HideConditions bool `yaml:"hide-conditions"`

	// HideHostname whether to hide the hostname in the Result
	HideHostname bool `yaml:"hide-hostname"`

	// HideURL whether to ensure the URL is not displayed in the results. Useful if the URL contains a token.
	HideURL bool `yaml:"hide-url"`

	// HidePort whether to hide the port in the Result
	HidePort bool `yaml:"hide-port"`

	// HideErrors whether to hide the errors in the Result
	HideErrors bool `yaml:"hide-errors"`

	// DontResolveFailedConditions whether to resolve failed conditions in the Result for display in the UI
	DontResolveFailedConditions bool `yaml:"dont-resolve-failed-conditions"`

	// ResolveSuccessfulConditions whether to resolve successful conditions in the Result for display in the UI
	ResolveSuccessfulConditions bool `yaml:"resolve-successful-conditions"`

	// Badge is the configuration for the badges generated
	Badge *Badge `yaml:"badge"`
}

type Badge struct {
	ResponseTime *ResponseTime `yaml:"response-time"`
}

type ResponseTime struct {
	Thresholds []int `yaml:"thresholds"`
}

var (
	ErrInvalidBadgeResponseTimeConfig = errors.New("invalid response time badge configuration: expected parameter 'response-time' to have 5 ascending numerical values")
)

// ValidateAndSetDefaults validates the UI configuration and sets the default values
func (config *Config) ValidateAndSetDefaults() error {
	if config.RecentChecksMaximumRows <= 0 {
		config.RecentChecksMaximumRows = 2
	}
	if config.RecentChecksResultsPerRow <= 0 {
		config.RecentChecksResultsPerRow = 50
	}
	if config.RecentChecksResultHeight == "" {
		config.RecentChecksResultHeight = "1.5rem"
	}
	if config.Badge != nil {
		if len(config.Badge.ResponseTime.Thresholds) != 5 {
			return ErrInvalidBadgeResponseTimeConfig
		}
		for i := 4; i > 0; i-- {
			if config.Badge.ResponseTime.Thresholds[i] < config.Badge.ResponseTime.Thresholds[i-1] {
				return ErrInvalidBadgeResponseTimeConfig
			}
		}
	} else {
		config.Badge = GetDefaultConfig().Badge
	}
	return nil
}

// GetDefaultConfig retrieves the default UI configuration
func GetDefaultConfig() *Config {
	return &Config{
		RecentChecksMaximumRows:     2,
		RecentChecksResultsPerRow:   50,
		RecentChecksResultHeight:    "1.5rem",
		HideHostname:                false,
		HideURL:                     false,
		HidePort:                    false,
		HideErrors:                  false,
		DontResolveFailedConditions: false,
		ResolveSuccessfulConditions: false,
		HideConditions:              false,
		Badge: &Badge{
			ResponseTime: &ResponseTime{
				Thresholds: []int{50, 200, 300, 500, 750},
			},
		},
	}
}
