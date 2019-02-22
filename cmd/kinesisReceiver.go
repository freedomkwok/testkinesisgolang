package main

import "github.com/sendgridlabs/go-kinesis/batchproducer"

type KinesisStatReceiver struct {
	PreviousStatsBatch *batchproducer.StatsBatch
	CurrentStatsbATCH  *batchproducer.StatsBatch
}

func NewKinesisStatReceiver() *KinesisStatReceiver {
	kinesisStatReceiver := new(KinesisStatReceiver)
	return kinesisStatReceiver
}

func (k *KinesisStatReceiver) Receive(statsBatch batchproducer.StatsBatch) {
	k.PreviousStatsBatch = k.CurrentStatsbATCH
	k.CurrentStatsbATCH = &statsBatch
}
