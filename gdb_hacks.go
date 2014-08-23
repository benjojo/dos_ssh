package main

import (
	"encoding/hex"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var gfb []byte
var UpdateScreenNow chan bool
var Pulling sync.Mutex

func StartPollingGDB() {
	UpdateScreenNow = make(chan bool)
	gfb = make([]byte, 0)
	nic, err := net.Dial("tcp", "localhost:1234")
	LazyHandle(err)
	for {
		select {
		case <-time.After(time.Second):
			Poll(nic)
		case <-UpdateScreenNow:
			Poll(nic)
		}
	}
}

func Poll(nic net.Conn) {
	Pulling.Lock()
	SendCMD(nic, "$g#67")
	for i := 0; i < 2; i++ {
		if i == 0 {
			SendCMD(nic, "$mb8000,800#5b") // BIOS Framebuffer ranges
		} else {
			SendCMD(nic, "$mb8800,7a0#93") // BIOS Framebuffer ranges
		}
		time.Sleep(time.Millisecond * 100) // You may be able to lower this
	}
	SendCMD(nic, "$k#6b")
	Pulling.Unlock()
}

func SendCMD(nic net.Conn, payload string) {
	buffer := make([]byte, 25565)

	_, err := nic.Write([]byte(payload))
	LazyHandle(err)
	in, err := nic.Read(buffer)
	LazyHandle(err)

	// Because I can't seem to figure out WHEN GDB is going to send stuff
	// I have to do what you are seeing below, Because the other commands
	// I am executing don't go above 1000 bytes output, I can presume that
	// anything above 1000 chars is the results of my memory dump. This
	// does mean however that we can get a out of order terminal, and that
	// does suck, but until I can figure out how to get a consistant output
	// it will have to stay like this.
	if in > 1000 {
		printtext(buffer, in)
	}

	_, err = nic.Write([]byte("+"))
	LazyHandle(err)
}

func LazyHandle(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var fbcount int = 0

func printtext(dump []byte, in int) {
	realdata := dump[2 : in-3]
	fuckit := strings.Split(string(realdata), "#")
	bin, err := hex.DecodeString(string(fuckit[0]))
	if err == nil {
		for i := 0; i < len(bin); i++ {
			gfb = append(gfb, bin[i])
		}
	}
	fbcount++
	if fbcount == 2 {
		fbcount = 0
		log.Println("Sent FB out")
		FrameBufferUpdate <- gfb
		gfb = []byte{}
	}

}
