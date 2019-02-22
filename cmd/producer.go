package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
	XID "github.com/rs/xid"
	Log "github.com/sirupsen/logrus"
	Lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var S3_LOG *Log.Logger

func main() {

}

func startImport() {
	awsRegionID := "us-west-2"
	awsSession, err := session.NewSession(&aws.Config{Region: &awsRegionID})
	if err != nil {
		panic(errors.Errorf("Fail To Retrieve Section"))
	}

	streamName := "teststream"
	S3_LOG = Log.New()
	S3_LOG.Out = &Lumberjack.Logger{
		Filename: "logfile",
		MaxSize:  20000,
		MaxAge:   20000,
	}

	println(awsSession, streamName)
	kinesisProducer := new(KinesisProducer)
	kinesisProducer.NewKinesisProducer(awsRegionID, streamName, S3_LOG)

	(*kinesisProducer.DefaultProducer).Start()

	for i := 0; i < 1000; i++ {
		partitionKey := fmt.Sprintf("%v_%v", XID.New().String(), "1")
		record := new(Record)
		record.ID = partitionKey
		dataByte, err := json.Marshal(record)
		if err != nil {
			println("error data: ", err)
		} else {
			println("Add data: ", record)
			kinesisProducer.Add(partitionKey, dataByte)
		}
	}
}
