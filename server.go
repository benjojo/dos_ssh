package main

import (
	"code.google.com/p/go.crypto/ssh"
	"github.com/mitchellh/go-vnc"
	"log"
	"net"
	"time"
)

var Keyin chan string

func main() {
	FrameBufferUpdate = make(chan []byte)
	Keyin = make(chan string, 100)
	FrameBufferSubscribers = make(map[string]chan []byte)
	go MessageHub(FrameBufferUpdate, FrameBufferSubscribers)

	log.Println("Starting GDB client")
	go StartPollingGDB()
	log.Println("Starting VNC client")
	go VNCKeyIn(Keyin)
	log.Println("Starting SSH server")
	StartSSH()
}

func ServeDOSTerm(channel ssh.Channel) {
	go ReadSSHIn(channel)
	MyID := randSeq(5)
	FBIN := make(chan []byte)
	FrameBufferSubscribers[MyID] = FBIN
	defer delete(FrameBufferSubscribers, MyID)
	FB := make([]byte, 0)
	for {
		FB = <-FBIN
		if len(FB) != 4000 {
			continue
		}
		channel.Write([]byte("\x1B[0;0H")) // <ESC>[2J
		log.Printf("DL: %d", len(FB))
		ptr := 0
		outbound := ""
		for {
			outbound = outbound + VESAtoVT100(FB[ptr+1])
			outbound = outbound + CorrectBadChars(FB[ptr])

			ptr = ptr + 2
			if ptr >= len(FB) {
				break
			}

		}
		_, err := channel.Write([]byte(outbound))
		if err != nil {
			return
		}
	}
}

func ReadSSHIn(channel ssh.Channel) {
	buffer := make([]byte, 2)
	for {
		_, err := channel.Read(buffer)
		if err != nil {
			return
		}

		Keyin <- string(buffer[0])

		time.Sleep(time.Millisecond * 10)
		UpdateScreenNow <- true
		time.Sleep(time.Millisecond * 200)
	}
}

func VNCKeyIn(Presses chan string) {
	vncnic, err := net.Dial("tcp", "localhost:5900")
	LazyHandle(err)

	vncconn, err := vnc.Client(vncnic, &vnc.ClientConfig{})
	LazyHandle(err)

	for in := range Keyin {
		Pulling.Lock()

		if in == "\r" || in == "\n" { // Enter
			vncconn.KeyEvent(uint32(0xFF0D), true)
			vncconn.KeyEvent(uint32(0xFF0D), false)
		} else if uint8([]byte(in)[0]) == 127 { // Backspace
			vncconn.KeyEvent(uint32(0xFF08), true)
			vncconn.KeyEvent(uint32(0xFF08), false)
		} else {
			vncconn.KeyEvent(uint32([]byte(in)[0]), true)
			vncconn.KeyEvent(uint32([]byte(in)[0]), false)
		}
		Pulling.Unlock()

	}

}
