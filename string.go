package arrays

type String struct{}

func (s *String) Contains(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (s *String) Remove(list []string, a string) []string {
	for i, v := range list {
		if v == a {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}
