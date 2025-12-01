package conv

type parseError string

func (e parseError) Error() string { return string(e) }

var (
	ErrSyntax = parseError("invalid syntax")
	ErrRange  = parseError("value out of range")
)

const intSize = 32 << (^uint(0) >> 63)

func isSpace(c byte) bool {
	switch c {
	case ' ', '\t', '\n', '\r', '\v', '\f':
		return true
	}
	return false
}

func trimSpace(s string) string {
	start := 0
	for start < len(s) {
		if isSpace(s[start]) {
			start++
			continue
		}
		break
	}
	end := len(s) - 1
	for end >= start {
		if isSpace(s[end]) {
			end--
			continue
		}
		break
	}
	return s[start : end+1]
}

func fastParseInt(s string, bitSize int) (int64, error, bool) {
	if len(s) == 0 {
		return 0, ErrSyntax, true
	}
	if isSpace(s[0]) || isSpace(s[len(s)-1]) {
		return 0, nil, false
	}

	maxLen := 0
	if bitSize == intSize {
		if intSize == 32 {
			maxLen = 10
		} else {
			maxLen = 19
		}
	} else {
		switch bitSize {
		case 64:
			maxLen = 19
		case 32:
			maxLen = 10
		default:
			return 0, nil, false
		}
	}
	if !(len(s) > 0 && len(s) < maxLen) {
		return 0, nil, false
	}

	i := 0
	neg := false
	if s[0] == '+' || s[0] == '-' {
		neg = s[0] == '-'
		i++
		if i == len(s) {
			return 0, ErrSyntax, true
		}
	}

	var n int64
	for ; i < len(s); i++ {
		d := s[i] - '0'
		if d > 9 {
			return 0, ErrSyntax, true
		}
		n = n*10 + int64(d)
	}
	if neg {
		n = -n
	}
	return n, nil, true
}

func parseSignedInt(s string, bitSize int) (int64, error) {
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
	if bitSize == 0 {
		bitSize = intSize
	}
	posMax := (uint64(1) << (bitSize - 1)) - 1
	negLimit := uint64(1) << (bitSize - 1)
	limit := posMax
	if neg {
		limit = negLimit
	}

	var acc uint64
	for ; i < len(v); i++ {
		c := v[i]
		if c < '0' || c > '9' {
			return 0, ErrSyntax
		}
		d := uint64(c - '0')
		if acc > (limit-d)/10 {
			return 0, ErrRange
		}
		acc = acc*10 + d
	}

	if neg {
		if acc > negLimit {
			return 0, ErrRange
		}
		return -int64(acc), nil
	}
	if acc > posMax {
		return 0, ErrRange
	}
	return int64(acc), nil
}

// 对外接口
func StringToInt8(s string) (int8, error) {
	v, err := parseSignedInt(s, 8)
	return int8(v), err
}

func StringToInt16(s string) (int16, error) {
	v, err := parseSignedInt(s, 16)
	return int16(v), err
}

func StringToInt32(s string) (int32, error) {
	if v, err, ok := fastParseInt(s, 32); ok {
		return int32(v), err
	}
	v, err := parseSignedInt(s, 32)
	return int32(v), err
}

func StringToInt64(s string) (int64, error) {
	if v, err, ok := fastParseInt(s, 64); ok {
		return v, err
	}
	return parseSignedInt(s, 64)
}

func StringToInt(s string) (int, error) {
	sLen := len(s)
	if intSize == 32 && (0 < sLen && sLen < 10) ||
		intSize == 64 && (0 < sLen && sLen < 19) {
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0, ErrSyntax
			}
		}
		n := 0
		for i := 0; i < len(s); i++ {
			ch := s[i] - '0'
			if ch > 9 {
				goto slow
			}
			n = n*10 + int(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n, nil
	}
slow:
	v, err := parseSignedInt(s, intSize)
	return int(v), err
}
