package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [username]")
		os.Exit(1)
	}

	username := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Kirim username sebagai pesan bergabung
	_, err = conn.Write([]byte("join:" + username))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Welcome to server ", username)

	go handleConn(conn)

	// Kirim pesan dari client ke server
	inputReader := bufio.NewReader(os.Stdin)
	for {
		message, _ := inputReader.ReadString('\n')
		message = strings.TrimSpace(message) // Hapus spasi di awal dan akhir
		if message == "exit" {
			_, err = conn.Write([]byte("left:exit"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			break
		}
		_, err = conn.Write([]byte("message:" + message))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("You have left the chat.") // Tambahkan log keluar
}

func handleConn(request *net.UDPConn) {
	for {
		var buf [512]byte
		n, _, err := request.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
