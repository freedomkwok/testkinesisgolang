package main

import (
	"os"
	"testing"
)

func TestKinesisProducer(t *testing.T) {

	os.Setenv("AWS_ACCESS_KEY_ID", "AKIAJRS3L52WK2KSX5VQ")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "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")
	startImport()
}
