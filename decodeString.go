package main

import (
	"strconv"
	"strings"
)

type stack struct {
	items []byte
}

func (s *stack) push(item byte) {
	s.items = append(s.items, item)
}

func (s *stack) pop() (byte, bool) {

	if len(s.items) == 0 {
		return 0, true
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func (s *stack) peek() byte {
	if len(s.items) == 0 {
		return 0
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	return item
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

func repeat(times int, s string) string {
	var sb strings.Builder
	for i := 0; i < times; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func solve(times int, aux *stack) {
	tmp := ""

	v, _ := aux.pop()
	for v != ']' {
		if v >= 'a' && v <= 'z' {
			tmp += string(v)
		}
		v, _ = aux.pop()
	}
	rep := repeat(times, tmp)
	for i := len(rep) - 1; i >= 0; i-- {
		aux.push(rep[i])
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func decodeString(s string) string {

	st := &stack{}
	aux := &stack{}
	ans := ""

	for _, c := range s {
		st.push(byte(c))
	}

	for !st.isEmpty() {
		cur, _ := st.pop()

		if isDigit(cur) {
			num := "" + string(cur)
			for {
				if !isDigit(st.peek()) || st.isEmpty() {
					break
				}
				cur, _ := st.pop()
				num += string(cur)
			}
			times, _ := strconv.ParseInt(reverseString(num), 10, 64)
			solve(int(times), aux)
		}

		if !isDigit(cur) {
			aux.push(cur)
		}
	}

	for !aux.isEmpty() {
		cur, _ := aux.pop()
		ans += string(cur)
	}

	return ans
}
