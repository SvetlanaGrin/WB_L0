package main

import (
	"WB_L0/publisher"
)

func main() {
	Sc := publisher.ConnectStan("nats-example")
	stop := make(chan bool)

	publisher.RunPublish(Sc, stop)
	<-stop
}
