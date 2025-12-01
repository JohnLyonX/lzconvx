package conv

const (
	maxFloat64Exp = 308
	minFloat64Exp = -324
)

const (
	maxPow10Pos = maxFloat64Exp
	maxPow10Neg = -minFloat64Exp
)

var (
	pow10Pos = make([]float64, maxPow10Pos+1)
	pow10Neg = make([]float64, maxPow10Neg+1)
)

func init() {
	pow10Pos[0] = 1
	for i := 1; i <= maxPow10Pos; i++ {
		pow10Pos[i] = pow10Pos[i-1] * 10
	}
	pow10Neg[0] = 1
	for i := 1; i <= maxPow10Neg; i++ {
		pow10Neg[i] = pow10Neg[i-1] / 10
	}
}

func LzFloat64(s string) (float64, error) {
	return parseFloat64(s)
}

func LzFloat32(s string) (float32, error) {
	v, err := parseFloat64(s)
	return float32(v), err
}

// Backward-compatible aliases.
func StringToFloat64(s string) (float64, error) { return LzFloat64(s) }
func StringToFloat32(s string) (float32, error) { return LzFloat32(s) }

func parseFloat64(s string) (float64, error) {
	v := trimSpace(s)
	if v == "" {
		return 0, ErrSyntax
	}

	i := 0
	neg := false
	if v[0] == '+' || v[0] == '-' {
		neg = v[0] == '-'
		i++
		if i == len(v) {
			return 0, ErrSyntax
		}
	}

	var mant uint64
	var digits int
	var fracDigits int
	seenDot := false
	digitSeen := false

	for i < len(v) {
		c := v[i]
		if c == '.' {
			if seenDot {
				return 0, ErrSyntax
			}
			seenDot = true
			i++
			continue
		}
		if c >= '0' && c <= '9' {
			digitSeen = true
			d := uint64(c - '0')
			if mant <= (1<<63-1)/10 {
				mant = mant*10 + d
			} else if mant <= (1<<63 - 1) {
				mant = mant
			}
			digits++
			if seenDot {
				fracDigits++
			}
			i++
			continue
		}
		break
	}

	if !digitSeen {
		return 0, ErrSyntax
	}

	exp := 0
	if i < len(v) && (v[i] == 'e' || v[i] == 'E') {
		i++
		if i == len(v) {
			return 0, ErrSyntax
		}
		expNeg := false
		if v[i] == '+' || v[i] == '-' {
			expNeg = v[i] == '-'
			i++
			if i == len(v) {
				return 0, ErrSyntax
			}
		}
		if v[i] < '0' || v[i] > '9' {
			return 0, ErrSyntax
		}
		for i < len(v) && v[i] >= '0' && v[i] <= '9' {
			exp = exp*10 + int(v[i]-'0')
			i++
			if exp > 1000 {
				break
			}
		}
		if expNeg {
			exp = -exp
		}
	}

	if i != len(v) {
		return 0, ErrSyntax
	}

	exp -= fracDigits

	if mant == 0 {
		if neg {
			return -0.0, nil
		}
		return 0, nil
	}
	if exp > maxFloat64Exp {
		return 0, ErrRange
	}
	if exp < minFloat64Exp {
		return 0, ErrRange
	}

	abs := float64(mant) * pow10(exp)
	if abs == 0 && exp < 0 {
		return 0, ErrRange
	}
	if neg {
		abs = -abs
	}
	if abs > 1e308 || abs < -1e308 {
		return 0, ErrRange
	}
	return abs, nil
}

func pow10(exp int) float64 {
	if exp >= 0 {
		if exp <= maxPow10Pos {
			return pow10Pos[exp]
		}
		// For completeness; parse already bounds-checks exp.
		return pow10Pos[maxPow10Pos] * pow10(exp-maxPow10Pos)
	}
	ne := -exp
	if ne <= maxPow10Neg {
		return pow10Neg[ne]
	}
	return pow10Neg[maxPow10Neg] * pow10(exp+maxPow10Neg)
}
