package main

import (
	"time"

	Kinesis "github.com/sendgridlabs/go-kinesis"
	"github.com/sendgridlabs/go-kinesis/batchproducer"

	Log "github.com/sirupsen/logrus"
)

type KinesisProducer struct {
	DefaultProducer *batchproducer.Producer
}

func (d *KinesisProducer) NewKinesisProducer(region string, streamName string, log *Log.Logger) *KinesisProducer {
	authCredential, newAuthErr := Kinesis.NewAuthFromMetadata()
	if newAuthErr != nil {
		println(newAuthErr)
	}

	kinesisClient := Kinesis.New(authCredential, region)
	kinesisStatReceiver := NewKinesisStatReceiver()
	config := batchproducer.Config{
		AddBlocksWhenBufferFull: true,
		BatchSize:               10,
		BufferSize:              500,
		FlushInterval:           1 * time.Second,
		Logger:                  log,
		MaxAttemptsPerRecord:    5,
		StatInterval:            1 * time.Second,
		StatReceiver:            kinesisStatReceiver}

	bp, batchProducerErr := batchproducer.New(kinesisClient, streamName, config)
	if batchProducerErr != nil {
		println(batchProducerErr)
	}

	d.DefaultProducer = &bp
	return d
}

func (d *KinesisProducer) Add(partitionKey string, data []byte) {
	err := (*d.DefaultProducer).Add(data, partitionKey)
	if err != nil {
		println("kinesisAdd Error: ", err)
	}
}
