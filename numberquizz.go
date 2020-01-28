package main

import (
    "fmt"
    "github.com/CaptainQuirk/numberquizz/generator"
    "bufio"
    "os"
    "strings"
    "flag"
)

func trimNl(text string) string {
    text = strings.TrimRight(text, "\n")

    return text
}

func main() {
    max := flag.Int("max", 15, "The maximum number to quizz for")
    flag.Parse()

    statement := generator.GenerateStatement(*max)
    solution := statement.Solve()

    reader := bufio.NewReader(os.Stdin)
    fmt.Print(statement.ToQuestion())
    text, _ := reader.ReadString('\n')
    answer := trimNl(text)

    if (strings.ToLower(answer) != solution) {
        fmt.Printf("Error ! Answer was « %s » and got « %s »\n", solution, answer)
    } else {
        fmt.Println("Congrats ! You win !")
    }
}
