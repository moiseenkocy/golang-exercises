package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    for i := 1; i < len(os.Args); i++ {
        fmt.Println(strconv.Itoa(i) + " " + os.Args[i])
    }
}
