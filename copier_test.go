package main

import (
	"bytes"
	"testing"
)

var source = []byte("1234567890")

func TestCopier_Next(t *testing.T) {
	reader := bytes.NewReader(source)
	writer := bytes.NewBuffer([]byte{})
	options := &Options{Offset: 2, Limit: 7, BlockSize: 3}

	copier, err := NewCopier(reader, writer, options)
	if err != nil {
		t.Errorf("Failed to initiate copier: %s", err)
	}

	// Only one step
	if err := copier.Next(); err != nil {
		t.Errorf("Failed to copy data: %s", err)
	}

	expect := string(source[options.Offset : options.Offset+options.BlockSize])
	if result := writer.String(); result != expect {
		t.Errorf("Invalid copied data: expect %s, got %s", expect, result)
	}

	// Full read
	for {
		if err := copier.Next(); err != nil {
			break
		}
	}
	expect = string(source[options.Offset : options.Offset+options.Limit])
	if result := writer.String(); result != expect {
		t.Errorf("Invalid full copied data: expect %s, got %s", expect, writer.String())
	}
}

func TestCopier_GetProgressPercent(t *testing.T) {
	reader := bytes.NewReader(source)
	writer := bytes.NewBuffer([]byte{})
	options := &Options{Offset: 2, Limit: 5, BlockSize: 3}

	copier, _ := NewCopier(reader, writer, options)
	_ = copier.Next()

	progress := copier.GetProgressPercent()
	expect := (float32(options.BlockSize) / (float32(options.Limit))) * 100
	if (expect-progress) < 0.00000001 && (progress-expect) < 0.00000001 {
		t.Errorf("Invalid copy progress: expect %.2f, got %.2f", expect, progress)
	}
}
