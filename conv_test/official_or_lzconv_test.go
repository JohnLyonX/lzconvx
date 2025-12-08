package convtest

import (
	"strconv"
	"testing"

	"github.com/JohnLyonX/lzconvx"
)

// ---------- int8 ------------

func BenchmarkLzInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzInt8("123")
	}
}

func BenchmarkStdInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseInt("123", 10, 8)
	}
}

// ---------- int16 ------------

func BenchmarkLzInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzInt16("12345")
	}
}

func BenchmarkStdInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseInt("12345", 10, 16)
	}
}

// ---------- int32 ------------

func BenchmarkLzInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzInt32("123456789")
	}
}

func BenchmarkStdInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseInt("123456789", 10, 32)
	}
}

// ---------- int64 ------------

func BenchmarkLzInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzInt64("123456789012345")
	}
}

func BenchmarkStdInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseInt("123456789012345", 10, 64)
	}
}

// ---------- int (Atoi 对比) ------------

func BenchmarkLzAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzAtoi("123456789")
	}
}

func BenchmarkStdAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.Atoi("123456789")
	}
}
