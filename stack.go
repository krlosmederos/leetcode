package main

type stackModel struct {
	items []int
}

func (s *stackModel) push(item int) {
	s.items = append(s.items, item)
}

func (s *stackModel) pop() (int, bool) {

	if len(s.items) == 0 {
		return 0, true
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func (s *stackModel) peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, true
	}
	return s.items[len(s.items)-1], true
}
