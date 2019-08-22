package main

import "errors"

type Options struct {
	SourceFile string
	ResultFile string
	Offset     int64
	Limit      int64
	BlockSize  int64
}

func NewOptions() *Options {
	return &Options{}
}

func (opt *Options) Validate() error {
	if opt.SourceFile == "" {
		return errors.New("path to source file is required")
	}

	if opt.ResultFile == "" {
		return errors.New("path to result file is required")
	}

	if opt.Offset < 0 {
		return errors.New("offset can not be less than zero")
	}

	if opt.Offset <= 0 {
		return errors.New("block size can not be equal or less than zero")
	}

	return nil
}
