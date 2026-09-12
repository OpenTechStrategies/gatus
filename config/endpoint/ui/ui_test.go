package ui

import (
	"errors"
	"testing"
)

func TestValidateAndSetDefaults(t *testing.T) {
	tests := []struct {
		name              string
		config            *Config
		wantErr           error
		wantMaximumRows   int
		wantResultsPerRow int
		wantResultHeight  string
	}{
		{
			name: "with-valid-config",
			config: &Config{
				RecentChecksMaximumRows: 4,
				Badge: &Badge{
					ResponseTime: &ResponseTime{Thresholds: []int{50, 200, 300, 500, 750}},
				},
			},
			wantMaximumRows:   4,
			wantResultsPerRow: 50,
			wantResultHeight:  "1.5rem",
			wantErr:           nil,
		},
		{
			name: "with-invalid-threshold-length",
			config: &Config{
				Badge: &Badge{
					ResponseTime: &ResponseTime{Thresholds: []int{50, 200, 300, 500}},
				},
			},
			wantErr:           ErrInvalidBadgeResponseTimeConfig,
			wantMaximumRows:   2,
			wantResultsPerRow: 50,
			wantResultHeight:  "1.5rem",
		},
		{
			name: "with-invalid-thresholds-order",
			config: &Config{
				Badge: &Badge{ResponseTime: &ResponseTime{Thresholds: []int{50, 200, 500, 300, 750}}},
			},
			wantErr:           ErrInvalidBadgeResponseTimeConfig,
			wantMaximumRows:   2,
			wantResultsPerRow: 50,
			wantResultHeight:  "1.5rem",
		},
		{
			name:              "with-no-badge-configured", // should give default badge cfg
			config:            &Config{},
			wantMaximumRows:   2,
			wantErr:           nil,
			wantResultsPerRow: 50,
			wantResultHeight:  "1.5rem",
		},
		{
			name:              "with-zero-maximum-rows",
			config:            &Config{RecentChecksMaximumRows: 0},
			wantMaximumRows:   2,
			wantErr:           nil,
			wantResultsPerRow: 50,
			wantResultHeight:  "1.5rem",
		},
		{
			name:              "with-negative-maximum-rows",
			config:            &Config{RecentChecksMaximumRows: -1},
			wantMaximumRows:   2,
			wantErr:           nil,
			wantResultsPerRow: 50,
			wantResultHeight:  "1.5rem",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.config.ValidateAndSetDefaults(); !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %v, got %v", tt.wantErr, err)
			}
			if tt.config.RecentChecksMaximumRows != tt.wantMaximumRows {
				t.Errorf("Expected RecentChecksMaximumRows %d, got %d", tt.wantMaximumRows, tt.config.RecentChecksMaximumRows)
			}
			if tt.config.RecentChecksResultsPerRow != tt.wantResultsPerRow {
				t.Errorf("Expected RecentChecksResultsPerRow %d, got %d", tt.wantResultsPerRow, tt.config.RecentChecksResultsPerRow)
			}
			if tt.config.RecentChecksResultHeight != tt.wantResultHeight {
				t.Errorf("Expected RecentChecksResultHeight %q, got %q", tt.wantResultHeight, tt.config.RecentChecksResultHeight)
			}
		})
	}
}

func TestGetDefaultConfig(t *testing.T) {
	config := GetDefaultConfig()
	if config.RecentChecksMaximumRows != 2 {
		t.Errorf("Expected default RecentChecksMaximumRows 2, got %d", config.RecentChecksMaximumRows)
	}
	if config.RecentChecksResultsPerRow != 50 {
		t.Errorf("Expected default RecentChecksResultsPerRow 50, got %d", config.RecentChecksResultsPerRow)
	}
	if config.RecentChecksResultHeight != "1.5rem" {
		t.Errorf("Expected default RecentChecksResultHeight 1.5rem, got %q", config.RecentChecksResultHeight)
	}
}
