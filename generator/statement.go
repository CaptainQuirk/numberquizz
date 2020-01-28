package generator

import (
    "fmt"
)

type Statement struct {
    number int
    representation string
}

func (statement Statement) ToQuestion() string {
    return fmt.Sprintf("Turn %d into %s: ", statement.number, statement.representation)
}
