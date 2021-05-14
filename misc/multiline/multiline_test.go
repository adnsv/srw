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

func TestCommonIndent(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"empty 1", []string{"", ""}, ""},
		{"none 1", []string{"a", ""}, ""},
		{"none 2", []string{" ", "a"}, ""},
		{"none 3", []string{"a", "b"}, ""},
		{"none 4", []string{" ", "a"}, ""},
		{"none 5", []string{" ", "\t"}, ""},
		{"none 6", []string{"\t", " "}, ""},
		{"same 1", []string{" a", " b"}, " "},
		{"same 2", []string{"\ta", "\tb"}, "\t"},
		{"same 3", []string{"\t a", "\t b"}, "\t "},
		{"same 4", []string{" \t a", " \t b"}, " \t "},
		{"shorter 1", []string{" \t a", " \t\tb"}, " \t"},
		{"shorter 2", []string{"    a", "   b"}, "   "},
		{"ingoring empty 1", []string{" a", ""}, " "},
		{"ingoring empty 2", []string{"\ta", ""}, "\t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonIndent(tt.args); got != tt.want {
				t.Errorf("CommonIndent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveCommonIndent(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{"none 1", []string{"a", " b"}, []string{"a", " b"}},
		{"none 2", []string{"\ta", " b"}, []string{"\ta", " b"}},
		{"simple 1", []string{" a", " b"}, []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.args
			RemoveCommonIndent(v)
			if !reflect.DeepEqual(v, tt.args) {
				t.Errorf("RemoveCommonIndent() failed")
			}
		})
	}
}
