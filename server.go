package main

import (
	"code.google.com/p/go.crypto/ssh"
	"github.com/mitchellh/go-vnc"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"time"
)

var GSigner ssh.Signer
var FrameBufferUpdate chan []byte
var FrameBufferSubscribers map[string]chan []byte
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

	PEM_KEY := LoadPrivKeyFromFile("./id_rsa")
	private, err := ssh.ParsePrivateKey(PEM_KEY)

	if err != nil {
		log.Fatal("Key failed to parse.")
	}
	GSigner = private

	SSHConfig := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			perms := ssh.Permissions{}
			return &perms, nil
		},
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			perms := ssh.Permissions{}
			return &perms, nil
		},
	}

	SSHConfig.AddHostKey(GSigner)

	listener, err := net.Listen("tcp", "0.0.0.0:2222")
	if err != nil {
		log.Fatalln("Could not start TCP listening on 0.0.0.0:2222")
	}
	log.Println("Waiting for TCP conns on 0.0.0.0:2222")

	for {
		nConn, err := listener.Accept()
		if err != nil {
			log.Println("WARNING - Failed to Accept TCP conn. RSN: %s / %s", err.Error(), err)
			continue
		}
		go HandleIncomingConn(nConn, SSHConfig)
	}
}

func TimeoutConnection(Done chan bool, nConn net.Conn) {
	select {
	case <-Done:
		return
	case <-time.After(time.Second * 10):
		nConn.Close()
	}
}

func HandleIncomingConn(nConn net.Conn, config *ssh.ServerConfig) {
	DoneCh := make(chan bool)
	go TimeoutConnection(DoneCh, nConn)
	_, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err == nil {
		DoneCh <- true
	}
	// Right now that we are out of annoying people land.

	defer nConn.Close()
	go DiscardRequests(reqs)

	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			log.Printf("WARNING - Rejecting %s Because they asked for a chan type %s that I don't have", nConn.RemoteAddr().String(), newChannel.ChannelType())
			continue
		}

		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Printf("WARNING - Was unable to Accept channel with %s", nConn.RemoteAddr().String())
			return
		}
		go DiscardRequests(requests)
		go ServeDOSTerm(channel)
	}

}

func DiscardRequests(in <-chan *ssh.Request) {
	for req := range in {
		// log.Printf("REQ: %s %s", req.Type, req.WantReply)
		if req.WantReply {
			if req.Type == "shell" || req.Type == "pty-req" {
				// log.Printf("REQ: %s %s '%s'", req.Type, req.WantReply, req.Payload)
				req.Reply(true, nil)
			} else {
				req.Reply(false, nil)
			}
		}
	}
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
		log.Printf("IN! %s", string(buffer[0]))

		Keyin <- string(buffer[0])

		time.Sleep(time.Millisecond * 10)
		UpdateScreenNow <- true
		time.Sleep(time.Millisecond * 200)
	}
}

func LoadPrivKeyFromFile(file string) []byte {
	privateBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("Failed to load private key")
	}
	return privateBytes
}

func MessageHub(Input chan []byte, Clients map[string]chan []byte) {

	for {
		inbound := <-Input
		for k, v := range Clients {
			log.Printf("Send payload to %s", k)
			v <- inbound
		}
	}

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func VNCKeyIn(Presses chan string) {
	vncnic, err := net.Dial("tcp", "localhost:5900")
	LazyHandle(err)

	vncconn, err := vnc.Client(vncnic, &vnc.ClientConfig{})
	LazyHandle(err)

	for in := range Keyin {
		Pulling.Lock()

		if in == "\r" || in == "\n" {
			vncconn.KeyEvent(uint32(0xFF0D), true)
			vncconn.KeyEvent(uint32(0xFF0D), false)
		} else if uint8([]byte(in)[0]) == 127 {
			vncconn.KeyEvent(uint32(0xFF08), true)
			vncconn.KeyEvent(uint32(0xFF08), false)
		} else {
			vncconn.KeyEvent(uint32([]byte(in)[0]), true)
			vncconn.KeyEvent(uint32([]byte(in)[0]), false)
		}
		Pulling.Unlock()

	}

}
