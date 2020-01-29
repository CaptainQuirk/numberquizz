package generator

import (
    "fmt"
    "strconv"
)

type Statement struct {
    number int
    representation Representation
}

func (statement Statement) Number() int {
    return statement.number
}

func (statement Statement) ToQuestion() string {
    return fmt.Sprintf("Turn %d into %s: ", statement.number, statement.representation.Name())
}

func (statement Statement) getRepresentationDecimalBase() int {
    return statement.representation.Base()
}

func (statement Statement) Solve() string {
    return strconv.FormatInt(int64(statement.number), statement.getRepresentationDecimalBase())
}
