package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func echo1() string {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    return s
}

func echo2() string {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    return s
}

func echo3() string {
    return strings.Join(os.Args[1:], " ")
}

func main() {
    repcount := 10000
    var s string

    // echo1
    start := time.Now()
    for c := 0; c < repcount; c++ {
        s = echo1()
    }
    fmt.Println(s)
    fmt.Println(time.Since(start).Seconds())

    // echo2
    start = time.Now()
    for c := 0; c < repcount; c++ {
        s = echo2()
    }
    fmt.Println(s)
    fmt.Println(time.Since(start).Seconds())

    // echo3
    start = time.Now()
    for c := 0; c < repcount; c++ {
        s = echo3()
    } 
    fmt.Println(s)
    fmt.Println(time.Since(start).Seconds())
}
