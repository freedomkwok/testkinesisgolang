package main

import (
	"fmt"
	"time"

	producer "github.com/a8m/kinesis-producer"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/pkg/errors"
	XID "github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

func main() {
	startImport()
}

func startImport() {
	awsRegionID := "us-west-2"
	awsSession, err := session.NewSession(&aws.Config{Region: &awsRegionID})
	if err != nil {
		panic(errors.Errorf("Fail To Retrieve Section"))
	}

	streamName := "teststream"
	log := logrus.New()

	client := kinesis.New(session.New(aws.NewConfig()))
	pr := producer.New(&producer.Config{
		StreamName:   "test",
		BacklogCount: 2000,
		Client:       client,
	})
	pr.Start()

	// Handle failures
	go func() {
		for r := range pr.NotifyFailures() {
			// r contains `Data`, `PartitionKey` and `Error()`
			log.Error(r)
		}
	}()

	go func() {
		for i := 0; i < 5000; i++ {
			partitionKey := fmt.Sprintf("%v_%v", XID.New().String(), "1")
			err := pr.Put([]byte("foo"), partitionKey)

			if err != nil {
				log.WithError(err).Fatal("error producing")
			}
		}
	}()

	time.Sleep(10 * time.Second)
	pr.Stop()
}
