package convtest

import (
	"testing"

	"liangzhanbo/lzconvx/conv"
)

func TestStringToFloat64(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    float64
		wantErr error
	}{
		{"simple int", "123", 123, nil},
		{"with sign", "-45", -45, nil},
		{"fraction", "3.14", 3.14, nil},
		{"leading/trailing space", "  2.5  ", 2.5, nil},
		{"exp positive", "1.5e2", 150, nil},
		{"exp negative", "12e-1", 1.2, nil},
		{"zero", "0", 0, nil},
		{"negative zero", "-0.0", -0.0, nil},
		{"double dot", "1.2.3", 0, conv.ErrSyntax},
		{"bad char", "1a2", 0, conv.ErrSyntax},
		{"empty", "   ", 0, conv.ErrSyntax},
		{"huge exp overflow", "1e309", 0, conv.ErrRange},
		{"tiny exp underflow", "1e-400", 0, conv.ErrRange},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StringToFloat64(tt.input)
			if err != tt.wantErr {
				t.Fatalf("input %q: want err %v, got %v", tt.input, tt.wantErr, err)
			}
			if tt.wantErr == nil && !floatEqual(got, tt.want) {
				t.Fatalf("input %q: want %v, got %v", tt.input, tt.want, got)
			}
		})
	}
}

func TestStringToFloat32(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    float32
		wantErr error
	}{
		{"basic", "12.5", 12.5, nil},
		{"exp", "7e1", 70, nil},
		{"overflow", "1e309", 0, conv.ErrRange},
		{"syntax", "abc", 0, conv.ErrSyntax},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.StringToFloat32(tt.input)
			if err != tt.wantErr {
				t.Fatalf("input %q: want err %v, got %v", tt.input, tt.wantErr, err)
			}
			if tt.wantErr == nil && !floatEqual(float64(got), float64(tt.want)) {
				t.Fatalf("input %q: want %v, got %v", tt.input, tt.want, got)
			}
		})
	}
}

func floatEqual(a, b float64) bool {
	if a == b {
		return true
	}
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	const eps = 1e-9
	return diff < eps
}
