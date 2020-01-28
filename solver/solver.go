package solver

import (
    "strconv"
    "github.com/CaptainQuirk/numberquizz/generator"
)

func Solve(statement generator.Statement) string {
    return strconv.FormatInt(int64(statement.GetNumber()), statement.GetRepresentationDecimalBase())
}
