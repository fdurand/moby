package jsonfilelog // import "github.com/fdurand/moby/daemon/logger/jsonfilelog"

import (
	"context"
	"encoding/json"
	"io"

	"github.com/fdurand/moby/api/types/backend"
	"github.com/fdurand/moby/daemon/logger"
	"github.com/fdurand/moby/daemon/logger/jsonfilelog/jsonlog"
	"github.com/fdurand/moby/daemon/logger/loggerutils"
	"github.com/fdurand/moby/pkg/tailfile"
)

// ReadLogs implements the logger's LogReader interface for the logs
// created by this driver.
func (l *JSONFileLogger) ReadLogs(config logger.ReadConfig) *logger.LogWatcher {
	return l.writer.ReadLogs(config)
}

func decodeLogLine(dec *json.Decoder, l *jsonlog.JSONLog) (*logger.Message, error) {
	l.Reset()
	if err := dec.Decode(l); err != nil {
		return nil, err
	}

	var attrs []backend.LogAttr
	if len(l.Attrs) != 0 {
		attrs = make([]backend.LogAttr, 0, len(l.Attrs))
		for k, v := range l.Attrs {
			attrs = append(attrs, backend.LogAttr{Key: k, Value: v})
		}
	}
	msg := &logger.Message{
		Source:    l.Stream,
		Timestamp: l.Created,
		Line:      []byte(l.Log),
		Attrs:     attrs,
	}
	return msg, nil
}

type decoder struct {
	rdr io.Reader
	dec *json.Decoder
	jl  *jsonlog.JSONLog
}

func (d *decoder) Reset(rdr io.Reader) {
	d.rdr = rdr
	d.dec = nil
	if d.jl != nil {
		d.jl.Reset()
	}
}

func (d *decoder) Close() {
	d.dec = nil
	d.rdr = nil
	d.jl = nil
}

func (d *decoder) Decode() (msg *logger.Message, err error) {
	if d.dec == nil {
		d.dec = json.NewDecoder(d.rdr)
	}
	if d.jl == nil {
		d.jl = &jsonlog.JSONLog{}
	}
	return decodeLogLine(d.dec, d.jl)
}

// decodeFunc is used to create a decoder for the log file reader
func decodeFunc(rdr io.Reader) loggerutils.Decoder {
	return &decoder{
		rdr: rdr,
		dec: nil,
		jl:  nil,
	}
}

func getTailReader(ctx context.Context, r loggerutils.SizeReaderAt, req int) (io.Reader, int, error) {
	return tailfile.NewTailReader(ctx, r, req)
}
