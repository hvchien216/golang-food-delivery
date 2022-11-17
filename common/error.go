package common

import (
	"errors"
	"log"
)

var (
	RecordNotFound = errors.New("record not found")
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("App recover", err)
	}
}
