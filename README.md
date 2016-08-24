# xmpp-latency
Tool that measures the latency of messages between two XMPP users

## How it works?
The tool is configured to use 2 different connections (to the same or different xmpp endpoints). The only requirement is that both users can communicate:
* One connection is used as the source of the communication and it will send a XMPP message with the current time.
* The other connection is used as the destination of the communication and will parse the received messages. If the message is a time object, it will calculate the time since the time object was generated.

## Build
```
$ go get
$ go build
```

## Example of configuration
```
---
source:
  jid: user1@xmpp-domain
  password: user1
  endpoint: xmpp-domain:5222
destination:
  jid: user2@xmpp-domain
  password: user2
  endpoint: xmpp-domain:5222
```

## Run
```
$ ./xmpp-latency
$ ./xmpp-latency -config /path/to/config.yaml
```
