package main

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"fmt"
	"log"
	"time"
)

func SendTime(destination *XMPPClient, source *xmpp.XMPP) {
	// Send the current time to the destination client
	msg := xmpp.Message{
		To: destination.JID,
		Body: fmt.Sprintf("%v", time.Now().Format(time.RFC3339Nano)),
	}
	source.Out <- msg
}

func SendPresence(con *xmpp.XMPP) {
	con.Out <- xmpp.Presence{}
}

func HandleMessages(con *xmpp.XMPP) {
	// Wait for messages and parse them
	select{
	case msg := <- con.In:
		err := parseMessage(msg)
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}
}

func parseMessage(msg interface{}) error {
	// Parse the given message.
	// If the body is a Time object, calculate difference between now and the object time

	switch v := msg.(type) {
	case error:
		log.Printf("error : %v\n", v)
		return v
	case *xmpp.Message:
		t, err := time.Parse(time.RFC3339Nano, v.Body)
		if err != nil {
			return err
		}
		log.Printf("The delivery took %v\n", time.Since(t))
	default:
		//ignore message
	}
	return nil
}