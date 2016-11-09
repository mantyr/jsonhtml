package main

import (
    "github.com/mantyr/jsonhtml"
    "flag"
    "fmt"
    "log"
)

func main() {
    flag.Parse()
    for _, file_address := range flag.Args() {
        s, err := jsonhtml.ConvertFile(file_address)

        if err != nil {
            log.Println("Convert error file", file_address, err.Error())
            continue
        }
        fmt.Println(s)
    }
}