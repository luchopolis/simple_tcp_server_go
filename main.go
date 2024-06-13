package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	lnserver, err := net.Listen("tcp", ":3000")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server running at 3000")
	for {
		conn, err := lnserver.Accept()
		if err != nil {
			fmt.Println(err)
			continue // jump to the other conection
		}
		go handleConnection(conn)
	}

}

func handleConnection(con net.Conn) {
	defer con.Close()

	reader := bufio.NewReader(con)
	fmt.Println(con.RemoteAddr().String())
	// keep listen for the client continous data
	// SEND BASIC COMMANDS
	con.Write([]byte("(editor) - Open Text Editor"))
	for {
		handleData, errorReader := reader.ReadString('\n')
		if errorReader != nil {
			if errorReader == io.EOF {
				fmt.Printf("A Client Close Connection:")
				return
			}
			fmt.Println(errorReader)
			return
		}

		// Print the incoming data
		fmt.Printf("Received: %s", handleData)
		fmt.Printf("%v", handleData)

		sanitizeIncomingData := strings.TrimSpace(handleData)
		if sanitizeIncomingData == "saludame" {
			con.Write([]byte("K Onda un Lijazo?\n"))
		}

		switch sanitizeIncomingData {
		case "editor":
			go func() {
				cmd := exec.Command("gnome-text-editor") // open the fedora editor
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
					con.Write([]byte("Error ejecutando Commando\n"))
				}
			}()

			con.Write([]byte("Done?\n"))
		}

	}

}
