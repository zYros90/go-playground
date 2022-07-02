package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	MQTT mqtt.Client
}

const (
	username = "guest"
	password = "guest"
	host     = "localhost:1883"
)

func NewClient() (*Client, error) {
	client := new(Client)
	url := fmt.Sprintf("tcp://%s", host)
	o := mqtt.NewClientOptions()
	o.AddBroker(url)
	o.SetClientID("go_mqtt_client")
	o.SetUsername(username)
	o.SetPassword(password)
	o.SetOnConnectHandler(client.connectHandler)
	o.SetConnectionLostHandler(client.connectLostHandler)
	o.SetReconnectingHandler(client.reconnectHandler)
	o.SetAutoReconnect(true)
	o.SetMaxReconnectInterval(5 * time.Second)
	mqc := mqtt.NewClient(o)
	token := mqc.Connect()
	if token.WaitTimeout(5*time.Second) && token.Error() != nil {
		return nil, token.Error()
	}
	client.MQTT = mqc
	return client, nil
}

func (c *Client) reconnectHandler(client mqtt.Client, opts *mqtt.ClientOptions) {
	fmt.Println("reconnecting...")
	time.Sleep(5 * time.Second)
}

func (c *Client) connectHandler(client mqtt.Client) {
	fmt.Println("connected")
}

func (c *Client) connectLostHandler(client mqtt.Client, err error) {
	fmt.Println("lost connection")
}

func main() {
	mc, err := NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	callback := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Println("received msg on topic: ", msg.Topic())
		fmt.Printf("msg: %s\n", msg.Payload())
	}
	token := mc.MQTT.Subscribe("/test", 0, callback)
	success := token.WaitTimeout(1 * time.Second)
	if !success {
		fmt.Println("timeout subscribing")
	}
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
