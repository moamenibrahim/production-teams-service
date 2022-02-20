package utils

import (
	"github.com/lithammer/shortuuid"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GenShortUUID() string {
	return shortuuid.New()
}
