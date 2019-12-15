package hello

import "testing"

func TestHoge(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"Hoge without name",
			"Hello, ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hoge(); got != tt.want {
				t.Errorf("Hoge() = %v, want %v", got, tt.want)
			}
		})
	}
}
