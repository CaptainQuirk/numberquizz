package main

import (
    "fmt"
    "github.com/CaptainQuirk/numberquizz/generator"
    "bufio"
    "os"
)

func main() {
    target := generator.GenerateTargetRepresentation()
    number := generator.GenerateNumberToTransform()

    question := fmt.Sprintf("Turn %d into %s: ", number, target)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print(question)
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}
