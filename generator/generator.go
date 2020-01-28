package generator

import (
    "math/rand"
    "time"
)

func GenerateStatement(max int) (statement Statement) {
    number := GenerateNumberToTransform(max)
    target := GenerateTargetRepresentation()

    return Statement{number, target}
}

func generateInt(max int) (number int) {
  rand.Seed(time.Now().UnixNano())
  number = rand.Intn(max)

  return
}

func GenerateNumberToTransform(max int) (int) {
    return generateInt(max)
}

func GenerateTargetRepresentation() (target string) {
    targets := [2]string{"binary", "hexadecimal"}
    index := generateInt(2)
    target = targets[index]

    return
}
