package generator

import (
    "math/rand"
    "time"
)

func GenerateNumberToTransform() (int) {
    return GenerateInt(255)
}

func GenerateInt(max int) (number int) {
    rand.Seed(time.Now().UnixNano())
    number = rand.Intn(max)

    return
}

func GenerateTargetRepresentation() (target string) {
    targets := [2]string{"binary", "hexadecimal"}
    index := GenerateInt(2)
    target = targets[index]

    return
}
