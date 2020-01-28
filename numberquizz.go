package main

import (
    "fmt"
    "github.com/CaptainQuirk/numberquizz/generator"
    "github.com/CaptainQuirk/numberquizz/solver"
    "bufio"
    "os"
    "strings"
)

func main() {
    statement := generator.GenerateStatement()
    solution := solver.Solve(statement)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print(statement.ToQuestion())
    text, _ := reader.ReadString('\n')

    if (strings.TrimRight(answer, "\n") != solution) {
        fmt.Printf("Error ! Answer was « %s » and got « %s »\n", solution, answer)
    } else {
        fmt.Println("Congrats ! You win !")
    }
}
