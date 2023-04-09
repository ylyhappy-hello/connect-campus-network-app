package apis

import "testing"

func TestGetQueryStringError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *GetQueryStringError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("GetQueryStringError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetQueryStringError_otherError(t *testing.T) {
	tests := []struct {
		name string
		e    *GetQueryStringError
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.otherError(); got != tt.want {
				t.Errorf("GetQueryStringError.otherError() = %v, want %v", got, tt.want)
			}
		})
	}
}
