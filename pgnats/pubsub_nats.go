package main

import (
	"fmt"
	"time"

	natssrv "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func main() {
	opts := &natssrv.Options{
		JetStream: true,
		StoreDir:  "data/jetstream",
	}
	ns, err := natssrv.NewServer(opts)
	if err != nil {
		panic(err)
	}
	go ns.Start()

	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

	nc, err := nats.Connect(ns.ClientURL())

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))

	if err != nil {
		panic(err)
	}

	// Simple Async Stream Publisher
	for i := 0; i < 500; i++ {
		js.PublishAsync("ORDERS.scratch", []byte(fmt.Sprintf("%d-msg", i)))
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}

	// // Simple Pull Consumer
	// sub, err := js.Sub("ORDERS.*", "MONITOR")
	// msgs, err := sub.Fetch(10)

	// for _, m := range msgs {
	// 	fmt.Printf("Received message: %s\n", string(m.Data))
	// }

	js.AddConsumer("ORDERS_CONSUMER", &nats.ConsumerConfig{}, nil)
	// Unsubscribe
	// sub.Unsubscribe()

	// Drain
	// sub.Drain()
	// subject := "my-subject"

	// curreMessages := int32(0)

	// nc.Subscribe(subject, func(msg *nats.Msg) {
	// 	newInt := atomic.AddInt32(&curreMessages, 1)
	// 	data := string(msg.Data)
	// 	fmt.Println(data)
	// 	if newInt == 22 {
	// 		ns.Shutdown()
	// 	}
	// })

	// for i := 0; i < 30; i++ {
	// 	nc.Publish(subject, []byte(fmt.Sprintf("Hello %d embedded NATS!", i)))
	// }

	// ns.WaitForShutdown()
}
