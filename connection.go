package main

import (
	"github.com/emgee/go-xmpp/src/xmpp"
)

func NewConnection(client *XMPPClient) (*xmpp.XMPP, error) {
	jid, err := xmpp.ParseJID(client.JID)
	if err != nil {
		return nil, err
	}

	stream, err := xmpp.NewStream(client.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	config := &xmpp.ClientConfig{NoTLS: true, InsecureSkipVerify: true}
	X, err := xmpp.NewClientXMPP(stream, jid, client.Password, config)
	if err != nil {
		return nil, err
	}

	return X, nil
}