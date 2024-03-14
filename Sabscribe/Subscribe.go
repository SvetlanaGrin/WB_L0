package Sabscribe

import (
	"WB_L0/internal/handler"
	"github.com/nats-io/stan.go"
	"log"
)

type Subscribers struct {
	handler *handler.Handler
}

func NewSubscribers(handler *handler.Handler) *Subscribers {
	return &Subscribers{handler: handler}
}

func ConnectStan(clientID string) stan.Conn {
	clusterID := "test-cluster"    // nats cluster id
	url := "nats://127.0.0.1:4222" // nats url
	// you can set client id anything
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url),
		stan.Pings(1, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, url)
	}

	log.Println("Connected Nats")

	return sc
}

func (s *Subscribers) RunSubscribers(Sc stan.Conn, stop chan bool) {
	ch := make(chan *stan.Msg)
	go func() {
		defer func() { stop <- true }()
		s.MessageResponder(Sc, "test", "test", "test-1", ch)
		for {
			select {
			case <-ch:
				a := <-ch
				s.handler.AddOrder(a.Data)
			case <-stop:
			}
		}
	}()
}

func (s *Subscribers) MessageResponder(Sc stan.Conn, subject, qgroup, durable string, ch chan *stan.Msg) {
	mcb := func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Printf("failed to ACK msg:%v", err)
		}
		ch <- msg
	}

	_, err := Sc.QueueSubscribe(subject,
		qgroup, mcb,
		stan.DeliverAllAvailable(),
		stan.SetManualAckMode(),
		stan.DurableName(durable))
	if err != nil {
		log.Fatalf("Not Subscribe", err)
	}
}
