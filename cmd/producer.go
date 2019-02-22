package main

import (
	"fmt"
	"time"

	producer "github.com/a8m/kinesis-producer"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	XID "github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"encoding/json"
	"math/rand"
	"strconv"
)

func main() {
	startImport()
}

func startImport() {
	awsRegionID := "us-west-2" 
	log := logrus.New()
	
	client := kinesis.New(session.New(&aws.Config{Region: &awsRegionID}))
	pr := producer.New(&producer.Config{
		StreamName:   "teststream2",//test
		BacklogCount: 2000,
		Client:       client,
		BatchCount: 10,
		AggregateBatchCount: 10,
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
			r := new(Recorder)
			r.ID = fmt.Sprintf("%v_%v", XID.New().String(), strconv.Itoa(i))
			r.Amount = rand.Intn(20)

			b , _ := json.Marshal(r)
			
			err := pr.Put(b, fmt.Sprintf("%v", rand.Intn(4)))

			if err != nil {
				log.WithError(err).Fatal("error producing")
			}
		}
	}()

	time.Sleep(30 * time.Second)
	pr.Stop()
}
