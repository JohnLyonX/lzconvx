package convtest

import (
	"testing"

	"github.com/JohnLyonX/lzconvx"
)

func TestStringToInt8(t *testing.T) {
	cases := []struct {
		input   string
		want    int8
		wantErr error
	}{
		{" 127 ", 127, nil},
		{"-128", -128, nil},
		{"128", 0, lzconvx.ErrRange},
		{"-129", 0, lzconvx.ErrRange},
		{"+", 0, lzconvx.ErrSyntax},
		{"12a", 0, lzconvx.ErrSyntax},
		{"   ", 0, lzconvx.ErrSyntax},
	}
	for _, tt := range cases {
		got, err := lzconvx.LzInt8(tt.input)
		if err != tt.wantErr {
			t.Fatalf("input %q: want err %v, got %v", tt.input, tt.wantErr, err)
		}
		if tt.wantErr == nil && got != tt.want {
			t.Fatalf("input %q: want %d, got %d", tt.input, tt.want, got)
		}
	}
}

func TestStringToInt16(t *testing.T) {
	cases := []struct {
		input   string
		want    int16
		wantErr error
	}{
		{"32767", 32767, nil},
		{"-32768", -32768, nil},
		{"32768", 0, lzconvx.ErrRange},
		{"-32769", 0, lzconvx.ErrRange},
		{"1 2", 0, lzconvx.ErrSyntax},
	}
	for _, tt := range cases {
		got, err := lzconvx.LzInt16(tt.input)
		if err != tt.wantErr {
			t.Fatalf("input %q: want err %v, got %v", tt.input, tt.wantErr, err)
		}
		if tt.wantErr == nil && got != tt.want {
			t.Fatalf("input %q: want %d, got %d", tt.input, tt.want, got)
		}
	}
}

func TestStringToInt32(t *testing.T) {
	cases := []struct {
		input   string
		want    int32
		wantErr error
	}{
		{"2147483647", 2147483647, nil},
		{"-2147483648", -2147483648, nil},
		{"2147483648", 0, lzconvx.ErrRange},
		{"-2147483649", 0, lzconvx.ErrRange},
	}
	for _, tt := range cases {
		got, err := lzconvx.LzInt32(tt.input)
		if err != tt.wantErr {
			t.Fatalf("input %q: want err %v, got %v", tt.input, tt.wantErr, err)
		}
		if tt.wantErr == nil && got != tt.want {
			t.Fatalf("input %q: want %d, got %d", tt.input, tt.want, got)
		}
	}
}

func TestStringToInt64(t *testing.T) {
	cases := []struct {
		input   string
		want    int64
		wantErr error
	}{
		{"9223372036854775807", 9223372036854775807, nil},
		{"-9223372036854775808", -9223372036854775808, nil},
		{"9223372036854775808", 0, lzconvx.ErrRange},
		{"-9223372036854775809", 0, lzconvx.ErrRange},
	}
	for _, tt := range cases {
		got, err := lzconvx.LzInt64(tt.input)
		if err != tt.wantErr {
			t.Fatalf("input %q: want err %v, got %v", tt.input, tt.wantErr, err)
		}
		if tt.wantErr == nil && got != tt.want {
			t.Fatalf("input %q: want %d, got %d", tt.input, tt.want, got)
		}
	}
}

func TestStringToInt(t *testing.T) {
	// 根据运行位宽选择边界
	is64 := (^uint(0) >> 63) == 1
	var maxStr, minStr, overflowStr string
	if is64 {
		maxStr = "9223372036854775807"
		minStr = "-9223372036854775808"
		overflowStr = "9223372036854775808"
	} else {
		maxStr = "2147483647"
		minStr = "-2147483648"
		overflowStr = "2147483648"
	}

	cases := []struct {
		input   string
		want    int
		wantErr error
	}{
		{" 42 ", 42, nil},
		{maxStr, intFromString(maxStr, is64), nil},
		{minStr, intFromString(minStr, is64), nil},
		{overflowStr, 0, lzconvx.ErrRange},
		{"-999999999999999999999", 0, lzconvx.ErrRange},
		{"abc", 0, lzconvx.ErrSyntax},
	}
	for _, tt := range cases {
		got, err := lzconvx.LzAtoi(tt.input)
		if err != tt.wantErr {
			t.Fatalf("input %q: want err %v, got %v", tt.input, tt.wantErr, err)
		}
		if tt.wantErr == nil && got != tt.want {
			t.Fatalf("input %q: want %d, got %d", tt.input, tt.want, got)
		}
	}
}

func intFromString(s string, is64 bool) int {
	// 仅用于将已知边界字符串转成 int，避免依赖其他库
	var neg bool
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '-' {
			neg = true
			continue
		}
		n = n*10 + int(c-'0')
	}
	if neg {
		n = -n
	}
	return n
}
