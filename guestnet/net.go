/**
 * Created by Administrator on 13-12-10.
 */
package main

import (
    "net"
    "fmt"
    "os"
)

/**
 * UDP 客户端
 */
func main() {

    pUDPAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7070")

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error ResolveUDPAddr: %s", err.Error())
        return
    }

    pUDPConn, err := net.DialUDP("udp", nil, pUDPAddr)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error DialUDP: %s", err.Error())
        return
    }

    n, err := pUDPConn.Write([]byte("你好啊！！！"))

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error WriteToUDP: %s", err.Error())
        return
    }

    fmt.Fprintf(os.Stdout, "writed: %d", n)

    buf := make([]byte, 1024)
    n, _, err = pUDPConn.ReadFromUDP(buf)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error ReadFromUDP: %s", err.Error())
        return
    }

    fmt.Fprintf(os.Stdout, "readed: %d  %s", n, string(buf[:n]))
}