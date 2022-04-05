package div

import "errors"

func Division(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir")
	}

	return a / b, nil
}
