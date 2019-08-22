package main

import (
	"io"
)

type Copier struct {
	Source    io.ReadSeeker
	Result    io.Writer
	Options   *Options
	BytesRead int64
}

func NewCopier(source io.ReadSeeker, result io.Writer, options *Options) (*Copier, error) {
	// Move inner file cursor
	if _, err := source.Seek(options.Offset, io.SeekStart); err != nil {
		return nil, err
	}

	return &Copier{source, result, options, 0}, nil
}

func (c *Copier) Next() error {
	// Change block size on last step
	bs := c.Options.BlockSize
	if c.Options.Limit > 0 && c.BytesRead+bs > c.Options.Limit {
		bs = c.Options.Limit - c.BytesRead
	}

	buf := make([]byte, bs)
	read, err := c.Source.Read(buf)
	c.BytesRead += int64(read)

	if read > 0 {
		if _, err := c.Result.Write(buf[:read]); err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	// Stop reading when the limit is reached
	if c.Options.Limit > 0 && c.BytesRead >= c.Options.Limit {
		return io.EOF
	}

	return nil
}

func (c *Copier) GetProgressPercent() float32 {
	return float32(c.BytesRead) / (float32(c.Options.Limit) / 100)
}
