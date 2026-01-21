package main

const (
	INT_MAX = 2147483647
	INT_MIN = -2147483648
)

func myAtoi(s string) int {
	n := len(s)
	i := 0

	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	sign := 1
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	var res int64 = 0
	for i < n {
		ch := s[i]
		if ch < '0' || ch > '9' {
			break
		}
		digit := int64(ch - '0')

		if res > int64(INT_MAX)/10 || (res == int64(INT_MAX)/10 && digit > int64(INT_MAX%10)) {
			if sign == 1 {
				return INT_MAX
			}
			return INT_MIN
		}

		res = res*10 + digit
		i++
	}

	val := int(res) * sign
	return val
}

