package test_repo

import (
	"fmt"
	"math/rand/v2"
)

func Core() string {
	i := rand.N(10000)
	return fmt.Sprintf("Core: %d", i)
}
