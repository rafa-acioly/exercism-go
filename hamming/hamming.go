package hamming

import "errors"

// Distance return a mutation quantity on given sequences
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New(
			`length of first argument and second argument must be equal`,
		)
	}

	var quantity int
	for index := range a {
		if a[index] != b[index] {
			quantity++
		}
	}

	return quantity, nil
}
