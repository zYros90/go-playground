package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	kafka "github.com/segmentio/kafka-go"
)

const (
	// to produce messages
	topic     = "my-topic"
	partition = 0
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	connWriter, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	connReader, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = writeMsgs(connWriter)
	if err != nil {
		log.Fatal(err)
	}

	err = consumeMsgs(connReader)
	if err != nil {
		log.Fatal(err)
	}

	<-sig
	err = connWriter.Close()
	if err != nil {
		log.Fatal("failed to close conn:", err)
	}

	err = connReader.Close()
	if err != nil {
		log.Fatal("failed to close conn:", err)
	}
}

func writeMsgs(conn *kafka.Conn) error {
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(
		kafka.Message{
			Topic: topic,
			Value: []byte("one!"),
		},
		kafka.Message{
			Topic: topic,
			Value: []byte("two!"),
		},
		kafka.Message{
			Topic: topic,
			Value: []byte("three!"),
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to write msgs")
	}

	return nil
}

func consumeMsgs(conn *kafka.Conn) error {
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	b := make([]byte, 10e3)            // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		return errors.Wrap(err, "failed to close batch")
	}
	return nil
}
