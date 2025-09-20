package repo

import (
	"github.com/google/uuid"
)

func Core() string {
	v7, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return v7.String()
}

func CoreV2() int {
	v7, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return int(v7.ID())
}
