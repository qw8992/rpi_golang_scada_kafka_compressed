package main

import (
	"fmt"
	"scada/uyeg"

	"github.com/Shopify/sarama"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("106.255.236.186:9012").Strings()
	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
	topic      = kingpin.Flag("topic", "Topic name").Default("Default").String()
)

func kafka(chValue chan string, client *uyeg.ModbusClient) {
	//kafka_Setting
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_0_1

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = *maxRetry
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionGZIP

	producer, err := sarama.NewSyncProducer([]string{
		"106.255.236.186:9011",
		"106.255.236.186:9012",
		"106.255.236.186:9013"}, config)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	//data get
	for {
		select {
		case values := <-chValue:
			messageSend(producer, values, client)
		}
	}
}

func messageSend(producer sarama.SyncProducer, values string, client *uyeg.ModbusClient) {
	*topic = fmt.Sprintf("m%s", client.Device.GatewayId)
	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.StringEncoder(values),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)
}
