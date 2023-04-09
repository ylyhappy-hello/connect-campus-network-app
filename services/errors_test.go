package services

import "testing"

func TestFormatCampusNetServiceError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *FormatCampusNetServiceError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("FormatCampusNetServiceError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCampusNetServiceError_otherError(t *testing.T) {
	tests := []struct {
		name string
		e    *FormatCampusNetServiceError
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.otherError(); got != tt.want {
				t.Errorf("FormatCampusNetServiceError.otherError() = %v, want %v", got, tt.want)
			}
		})
	}
}
