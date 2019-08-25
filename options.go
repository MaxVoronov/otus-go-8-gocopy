package main

import "errors"

type Options struct {
	SourceFile string
	ResultFile string
	Offset     int64
	Limit      int64
	BlockSize  int64
}

var ErrorSourcePathEmpty = errors.New("path to source file is required")
var ErrorResultPathEmpty = errors.New("path to result file is required")
var ErrorOffsetLessZero = errors.New("offset can not be less than zero")
var ErrorBlockSizeLessZero = errors.New("block size can not be equal or less than zero")

func NewOptions() *Options {
	return &Options{}
}

func (opt *Options) Validate() error {
	if opt.SourceFile == "" {
		return ErrorSourcePathEmpty
	}

	if opt.ResultFile == "" {
		return ErrorResultPathEmpty
	}

	if opt.Offset < 0 {
		return ErrorOffsetLessZero
	}

	if opt.BlockSize <= 0 {
		return ErrorBlockSizeLessZero
	}

	return nil
}
