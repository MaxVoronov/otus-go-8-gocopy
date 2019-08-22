package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// gocopy ­from /path/to/source ­to /path/to/dest ­offset 1024 ­limit 4096 -bs 1024
	options := NewOptions()
	flag.StringVar(&options.SourceFile, "from", "", "Path to source file")
	flag.StringVar(&options.ResultFile, "to", "", "Path to result file")
	flag.Int64Var(&options.Offset, "offset", 0, "Offset from start of file")
	flag.Int64Var(&options.Limit, "limit", 0, "Read bytes limit")
	flag.Int64Var(&options.BlockSize, "bs", 1024, "Read/write block size (bytes)")
	flag.Parse()

	if err := options.Validate(); err != nil {
		log.Fatalf("Error: %s", err)
	}

	if err := process(os.Stdout, options); err != nil {
		log.Fatalf("Processing error: %s", err)
	}
}

func process(out io.Writer, options *Options) error {
	fmt.Fprintf(out,
		"Copy \"%s\" => \"%s\" [offset: %d; limit: %d; bs %d]\n",
		options.SourceFile,
		options.ResultFile,
		options.Offset,
		options.Limit,
		options.BlockSize,
	)

	src, err := os.Open(options.SourceFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, _ := os.Create(options.ResultFile)
	defer dst.Close()

	// Update limit for showing correct progress data
	srcInfo, err := src.Stat()
	if err != nil {
		return err
	}
	if options.Limit == 0 || options.Limit > srcInfo.Size()-options.Offset {
		options.Limit = srcInfo.Size() - options.Offset
	}

	copier, err := NewCopier(src, dst, options)
	if err != nil {
		return err
	}

	for {
		err := copier.Next()
		fmt.Fprintf(out,
			"Processing... [%d / %d] %.1f%%\r",
			copier.BytesRead,
			copier.Options.Limit,
			copier.GetProgressPercent(),
		)

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	fmt.Fprintf(out, "\nDone!\n")

	return nil
}
