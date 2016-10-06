package main

import (
    "fmt"
    "flag"
)

func main() {
    fmt.Println("Hola wurld!")

    outFile := flag.String("o", "output.txt", "an output file")

    flag.Parse()

    fmt.Println(*outFile)
}
