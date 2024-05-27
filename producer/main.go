package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// (1) Configuração e Criação do Produtor
	config := &kafka.ConfigMap{"bootstrap.servers": "localhost"}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	// (1) FIM

	// (2) Produção da Mensagem
	topic := "anthony"
	msg := &kafka.Message{
		Value:          []byte("helloworld"),
		Key:            []byte("1"),
		Headers: []kafka.Header{
			{Key: "regiao", Value: []byte("SC")},
		},
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	}

	wait := make(chan kafka.Event)
	err = producer.Produce(msg, wait)
	if err != nil {
		panic(err)
	}
	// (2) FIM

	<-wait
}
