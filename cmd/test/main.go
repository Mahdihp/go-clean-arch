package main

import (
	"encoding/json"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/bxcodec/go-clean-arch/internal/bybit_ws/params"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	inputOrderBook := params.InputOrderBook{
		Op:   "subscribe",
		Args: []string{"orderbook.200.BTCUSDT"},
	}
	output := make(chan bool, 1)
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               10,
		MaxConcurrentRequests: 5,
		ErrorPercentThreshold: 5,
	})
	errors := hystrix.Go("my_command", func() error {
		// talk to other services
		err := getOrderBook(inputOrderBook)
		return err
	}, func(err error) error {
		// do this when services are down
		fmt.Println(err)
		return nil
	})
	for {
		select {
		case <-output:
			// success
			fmt.Println("success")

		case <-errors:
			// failure
		}
	}
}

func getOrderBook(iob params.InputOrderBook) error {

	inputMsg, err2 := json.Marshal(iob)
	if err2 != nil {
		log.Println("Marshal err:", err2)
	}
	wsConn, _, err := websocket.DefaultDialer.Dial("wss://stream.bybit.com/v5/public/spot", nil)
	//err = wsConn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	if err != nil {
		return err
	}
	defer wsConn.Close()
	if err != nil {
		log.Fatal("Dial:", err)
		return err
	}
	err2 = wsConn.WriteMessage(websocket.TextMessage, inputMsg)
	if err2 != nil {
		log.Fatal("WriteMessage:", err2)
		return err2
	}
	var snapShot params.SpotOrderBook
	for {
		readData(wsConn, snapShot)
	}
	return nil
}
func readData(wsc *websocket.Conn, snapShot params.SpotOrderBook) error {
	mt, message, err := wsc.ReadMessage()
	if err != nil {
		log.Println("read :", err, mt)
		if mt == -1 {
		}
		return err
	}
	//log.Printf("recv: %s", message)
	if len(message) > 0 {
		var delta params.SpotOrderBook
		err = json.Unmarshal(message, &delta)
		if err != nil {
			log.Println("unmarshal err:", err)
		}
		if delta.Type == "snapshot" {
			log.Printf("recv: %s\n", message)
			snapShot = delta
		} else if delta.Type == "delta" {
			snapShot.UpdateSnapShot(delta.Data, 1000)
			log.Printf("recv: %d %d \n", len(snapShot.Data.A), len(snapShot.Data.B))
			//marshal, _ := json.Marshal(delta)
			//log.Printf("recv: %s", marshal)

		}
	}
	return nil
}
