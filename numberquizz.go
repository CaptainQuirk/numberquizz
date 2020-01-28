package main

import (
    "fmt"
    "github.com/CaptainQuirk/numberquizz/generator"
    "bufio"
    "os"
)

func main() {
    statement := generator.GenerateStatement()

    reader := bufio.NewReader(os.Stdin)
    fmt.Print(statement.ToQuestion())
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}
