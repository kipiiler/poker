package main

import (
    "fmt"
    "net"
    "bufio"
    "strings"
    "os"

)


func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
        //os.exit(1)
    }
    defer listener.Close()
    fmt.println("Listening on port 8080")

    for {

}
