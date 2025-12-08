package convtest

import (
	"strconv"
	"testing"

	"github.com/JohnLyonX/lzconvx"
)

// --- 测试用例组 ---
var (
	smallFloat     = "123.456"
	largeFloat     = "12345678901234567890.987654321"
	sciFloat1      = "1.23456e10"
	sciFloatNegExp = "9.87654e-12"
	tinyFloat      = "0.000000000123456"
	justIntFloat   = "42"
)

// ---------------- LzFloat64 vs ParseFloat ----------------

func BenchmarkLzFloat64_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzFloat64(smallFloat)
	}
}

func BenchmarkStdFloat64_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseFloat(smallFloat, 64)
	}
}

func BenchmarkLzFloat64_Sci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzFloat64(sciFloat1)
	}
}

func BenchmarkStdFloat64_Sci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseFloat(sciFloat1, 64)
	}
}

func BenchmarkLzFloat64_NegExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzFloat64(sciFloatNegExp)
	}
}

func BenchmarkStdFloat64_NegExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseFloat(sciFloatNegExp, 64)
	}
}

func BenchmarkLzFloat64_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzFloat64(largeFloat)
	}
}

func BenchmarkStdFloat64_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseFloat(largeFloat, 64)
	}
}

// ---------------- LzFloat32 vs ParseFloat ----------------

func BenchmarkLzFloat32_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzFloat32(smallFloat)
	}
}

func BenchmarkStdFloat32_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseFloat(smallFloat, 32)
	}
}

func BenchmarkLzFloat32_Sci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lzconvx.LzFloat32(sciFloat1)
	}
}

func BenchmarkStdFloat32_Sci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.ParseFloat(sciFloat1, 32)
	}
}
