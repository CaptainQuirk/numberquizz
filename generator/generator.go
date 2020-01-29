package generator

import (
    "math/rand"
    "time"
)

func GenerateStatement(max int) (statement Statement) {
    number := generateNumberToTransform(max)
    representation := generateTargetRepresentation()

    return Statement{number, representation}
}

func generateInt(max int) (number int) {
  rand.Seed(time.Now().UnixNano())
  number = rand.Intn(max)

  return
}

func generateNumberToTransform(max int) (int) {
    return generateInt(max)
}

func generateTargetRepresentation() (representation Representation) {
    var representations []Representation

    representations = append(representations, Representation{"binary", 10})
    representations = append(representations, Representation{"hexadecimal", 16})

    index := generateInt(2)
    representation = representations[index]

    return
}
