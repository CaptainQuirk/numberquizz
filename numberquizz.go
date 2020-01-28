package main

import (
    "fmt"
    "github.com/CaptainQuirk/numberquizz/generator"
)

func main() {
    question := fmt.Sprintf("Turn %d into %s: ", generator.GenerateInt(255), generator.GenerateTarget())

    fmt.Println(question)
}
