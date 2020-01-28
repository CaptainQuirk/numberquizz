package generator

import (
    "math/rand"
    "time"
)

func GenerateInt(max int) (number int) {
    rand.Seed(time.Now().UnixNano())
    number = rand.Intn(max)

    return
}
