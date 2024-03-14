package publisher

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func UUID() string {
	// Create a Version 4 UUID.
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return u2.String()
}
func RunPublish(Sc stan.Conn, stop chan bool) {
	ticker := time.NewTicker(time.Second)
	go func() {
		defer func() { stop <- true }()
		order := InitOrder()
		for {
			select {
			case <-ticker.C:
				order.OrderUid = UUID()
				dataOrder, err := json.Marshal(order)
				if err != nil {
					return
				}
				data := dataOrder
				PublishNats(Sc, data, "test")

			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

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

func PublishNats(Sc stan.Conn, data []byte, channel string) {
	ach := func(s string, err2 error) {}
	_, err := Sc.PublishAsync(channel, data, ach)
	if err != nil {
		log.Fatalf("Error during async publish: %v\n", err)
	}
}
