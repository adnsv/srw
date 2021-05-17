package multiline

import (
	"reflect"
	"testing"
)

func TestSplitAtEOLs(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{"empty", "", []string{}},
		{"single", "ab", []string{"ab"}},
		{"pair1", "a\nb", []string{"a", "b"}},
		{"pair2", "a\rb", []string{"a", "b"}},
		{"pair3", "a\r\nb", []string{"a", "b"}},
		{"pair4", "a\n\rb", []string{"a", "", "b"}},
		{"edge1", "\n", []string{"", ""}},
		{"edge2", "a\n", []string{"a", ""}},
		{"edge3", "\nb", []string{"", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitAtEOLs(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitAtEOLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
