package slices

type String struct{}

func (s String) Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func (s String) Remove(slice []string, str string) []string {
	for i, v := range slice {
		if v == str {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (s String) Equals(sliceA []string, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for i := range sliceA {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}
