package util

import (
	"regexp"
	"testing"
)

func TestGetDefaultDownloadPath(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		wantErr bool
	}{
		{
			name:    "Sucessfull for windows",
			pattern: `^C:\\Users\\[^\\]+\\Downloads$`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDefaultDownloadPath()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDefaultDownloadPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			matched, _ := regexp.MatchString(tt.pattern, got)
			if !matched {
				t.Errorf("GetDefaultDownloadPath() = %v, does not match pattern %v", got, tt.pattern)
			}
		})
	}
}
