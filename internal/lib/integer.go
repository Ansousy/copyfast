package lib

import "errors"

func Division(a int, b int) (error, float64) {

	if b == 0 {
		return errors.New("Division par 0 impossible"), 0
	}

	return nil, float64(a) / float64(b)
}
