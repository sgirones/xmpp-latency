package main

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"time"
	"flag"
)

func main() {
	var configPath = flag.String("config", "./config.yaml", "Path to the config file")
	flag.Parse()

	config, err := LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	sourceCon, err := NewConnection(config.Source)
	if err != nil {
		panic(err)
	}
	sourceCon.Out <- xmpp.Presence{}

	destinationCon, err := NewConnection(config.Destination)
	if err != nil {
		panic(err)
	}
	destinationCon.Out <- xmpp.Presence{}

	// Send current time to client2 from client1 every second
	go func() {
		for {
			SendTime(config.Destination, sourceCon)
			time.Sleep(1 * time.Second)
		}
	}()

	// Send presence messages every 10 seconds
	go func() {
		for {
			SendPresence(sourceCon)
			SendPresence(destinationCon)
			time.Sleep(10 * time.Second)
		}
	}()

	// Handle all messages received on con2
	for {
		HandleMessages(destinationCon)
	}
}