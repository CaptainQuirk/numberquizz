package main

import (
    "fmt"
    "github.com/CaptainQuirk/numberquizz/generator"
    "bufio"
    "os"
)

func main() {
    question := fmt.Sprintf("Turn %d into %s: ", generator.GenerateInt(255), generator.GenerateTarget())

    reader := bufio.NewReader(os.Stdin)
    fmt.Print(question)
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}
