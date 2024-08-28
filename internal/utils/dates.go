package utils

import "time"

func IsValidDate(d string) bool {
	_, err := time.Parse(time.DateOnly, d)

	return err == nil
}
