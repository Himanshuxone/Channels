package main

import(
	"fmt"
	"net"
	"time"
)

func sendMsg(msg, addr string) error {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        return err
    }
    defer conn.Close()
    _, err = fmt.Fprint(conn, msg)
    return err
} 

func broadcastMsg(msg string, addrs []string) error {
    errc := make(chan error)
    for _, addr := range addrs {
        go func(addr string) {
            errc <- sendMsg(msg, addr)
            fmt.Println("done")
        }(addr)
    }

    for _ = range addrs {
        if err := <-errc; err != nil {
            fmt.Println("Error returned from client channel")
            fmt.Println(err)
            return err
        }
    }
    return nil
}

func main() {
    addr := []string{"localhost:8080"}
    err := broadcastMsg("hi", addr)

    time.Sleep(time.Second)

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("everything went fine")
}