package generator

import (
    "math/rand"
    "time"
)

func GenerateStatement() (statement Statement) {
    number := GenerateNumberToTransform()
    target := GenerateTargetRepresentation()

    return Statement{number, target}
}

func generateInt(max int) (number int) {
  rand.Seed(time.Now().UnixNano())
  number = rand.Intn(max)

  return
}

func GenerateNumberToTransform() (int) {
    return generateInt(255)
}

func GenerateTargetRepresentation() (target string) {
    targets := [2]string{"binary", "hexadecimal"}
    index := generateInt(2)
    target = targets[index]

    return
}
