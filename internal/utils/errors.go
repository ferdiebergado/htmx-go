package utils

import (
	"fmt"
	"log"
)

func LogError(desc string, err error) {
	details := fmt.Errorf("error: %v", err)
	log.Printf("ERROR: %s %s", desc, details)
}
