package main

import (
	"code.google.com/p/go.crypto/ssh"
	"log"
	"net"
	"time"
)

var FrameBufferUpdate chan []byte
var FrameBufferSubscribers map[string]chan []byte

// Start listening for SSH connections
func StartSSH() {
	PEM_KEY := LoadPrivKeyFromFile("./id_rsa")
	private, err := ssh.ParsePrivateKey(PEM_KEY)
	if err != nil {
		log.Fatal("Key failed to parse.")
	}

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

	SSHConfig.AddHostKey(private)

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
		go HandleIncomingSSHConn(nConn, SSHConfig)
	}
}

// Wait 10 seconds before closing the connection (To stop dead connections)
func TimeoutConnection(Done chan bool, nConn net.Conn) {
	select {
	case <-Done:
		return
	case <-time.After(time.Second * 10):
		nConn.Close()
	}
}

func HandleIncomingSSHConn(nConn net.Conn, config *ssh.ServerConfig) {
	DoneCh := make(chan bool)
	go TimeoutConnection(DoneCh, nConn)
	_, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err == nil {
		DoneCh <- true
	}
	// Right now that we are out of annoying people land.

	defer nConn.Close()
	go HandleSSHrequests(reqs)

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
		go HandleSSHrequests(requests)
		go ServeDOSTerm(channel)
	}

}

func HandleSSHrequests(in <-chan *ssh.Request) {
	for req := range in {
		if req.WantReply {
			// Ensure that the other end does not panic that we don't offer terminals
			if req.Type == "shell" || req.Type == "pty-req" {
				req.Reply(true, nil)
			} else {
				req.Reply(false, nil)
			}
		}
	}
}
