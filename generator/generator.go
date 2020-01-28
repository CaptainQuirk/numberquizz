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

func GenerateTarget() (target string) {
    targets := [2]string{"binary", "hexadecimal"}
    index := GenerateInt(2)
    target = targets[index]

    return
}
