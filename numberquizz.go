package main

import (
    "fmt"
    "github.com/CaptainQuirk/numberquizz/generator"
    "github.com/CaptainQuirk/numberquizz/solver"
    "bufio"
    "os"
    "strings"
)

func filter(text string) string {
    text = strings.TrimRight(text, "\n")
    text = strings.ToLower(text)

    return text
}

func main() {
    statement := generator.GenerateStatement()
    solution := solver.Solve(statement)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print(statement.ToQuestion())
    text, _ := reader.ReadString('\n')
    answer := filter(text)

    if (answer != solution) {
        fmt.Printf("Error ! Answer was « %s » and got « %s »\n", solution, answer)
    } else {
        fmt.Println("Congrats ! You win !")
    }
}
