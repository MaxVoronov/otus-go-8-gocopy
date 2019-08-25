package main

import "testing"

func TestOptions_Validate(t *testing.T) {
	options := NewOptions()

	if err := options.Validate(); err == nil {
		t.Error("Should fail validation")
	}

	options.SourceFile = "path/to/source_file"
	if err := options.Validate(); err == ErrorSourcePathEmpty {
		t.Error("Wrong validation error (expect ErrorSourcePathEmpty)")
	}

	options.ResultFile = "path/to/result_file"
	if err := options.Validate(); err == ErrorResultPathEmpty {
		t.Error("Wrong validation error (expect ErrorSourcePathEmpty)")
	}

	options.Offset = 0
	if err := options.Validate(); err == ErrorOffsetLessZero {
		t.Error("Wrong validation error (expect ErrorOffsetLessZero)")
	}

	options.BlockSize = 1024
	if err := options.Validate(); err == ErrorBlockSizeLessZero {
		t.Error("Wrong validation error (expect ErrorBlockSizeLessZero)")
	}
}
