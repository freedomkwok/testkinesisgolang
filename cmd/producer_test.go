package main

import (
	"os"
	"testing"
)

func TestKinesisProducer(t *testing.T) {

	os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIPN3ACS5L6LR7ZYQ")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "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")
	startImport()
}
