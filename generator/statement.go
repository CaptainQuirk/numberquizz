package generator

import (
    "fmt"
    "strconv"
)

type Statement struct {
    number int
    representation string
}

func (statement Statement) GetNumber() int {
    return statement.number
}

func (statement Statement) ToQuestion() string {
    return fmt.Sprintf("Turn %d into %s: ", statement.number, statement.representation)
}

func (statement Statement) getRepresentationDecimalBase() int {
    var decimal_base int
    switch (statement.representation) {
    case "binary":
      decimal_base = 2
    case "hexadecimal":
      decimal_base = 16
    }

    return decimal_base
}

func (statement Statement) Solve() string {
    return strconv.FormatInt(int64(statement.number), statement.getRepresentationDecimalBase())
}
