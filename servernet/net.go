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
 * UDP 服务器
 */
func main() {
	pUDPAddr, err := net.ResolveUDPAddr("udp", ":7070")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	pUDPConn, err := net.ListenUDP("udp", pUDPAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	defer pUDPConn.Close()

	for {

		buf := make([]byte, 256)
		n , pUDPAddr, err := pUDPConn.ReadFromUDP(buf)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
		fmt.Fprintf(os.Stdout, "readed: %d", n)

		n, err = pUDPConn.WriteToUDP(buf, pUDPAddr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
		fmt.Fprintf(os.Stdout, "writed: %d", n)
	}
}