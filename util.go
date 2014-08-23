package main

import (
	"io/ioutil"
	"log"
	"math/rand"
)

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
