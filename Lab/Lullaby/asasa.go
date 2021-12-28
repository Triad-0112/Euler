package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

func getConfig() zap.Config {
	// your current config options
	return zap.NewProductionConfig()
}

type prependEncoder struct {
	// embed a zapcore encoder
	// this makes prependEncoder implement the interface without extra work
	zapcore.Encoder

	// zap buffer pool
	pool buffer.Pool
}

// EncodeEntry implementing only EncodeEntry
func (e *prependEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// new log buffer
	buf := e.pool.Get()

	// prepend the JournalD prefix based on the entry level
	buf.AppendString(e.toJournaldPrefix(entry.Level))
	buf.AppendString(" ")

	// calling the embedded encoder's EncodeEntry to keep the original encoding format
	consolebuf, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	// just write the output into your own buffer
	_, err = buf.Write(consolebuf.Bytes())
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// some mapper function
func (e *prependEncoder) toJournaldPrefix(lvl zapcore.Level) string {
	switch lvl {
	case zapcore.DebugLevel:
		return "<7>"
	case zapcore.InfoLevel:
		return "<6>"
	case zapcore.WarnLevel:
		return "<4>"
	}
	return ""
}

func main() {
	cfg := getConfig()

	// constructing our prependEncoder with a ConsoleEncoder using your original configs
	enc := &prependEncoder{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		pool:    buffer.NewPool(),
	}

	logger := zap.New(
		zapcore.NewCore(
			enc,
			os.Stdout,
			zapcore.DebugLevel,
		),
		// this mimics the behavior of NewProductionConfig.Build
		zap.ErrorOutput(os.Stderr),
	)
	//a := logger.With(zap.String("bar", "addefrg"))
	logger.With(zap.String("for", "baa")).Info("A")
	a := zap.String("for", "bar")
	logger.Info("this is info")
	logger.Debug("this is debug")
	logger.Warn("this is warn")
	logger.Info("The prefix will show " + a.String)
	logger.With(zap.String("sdd", "foo")).Info("SSS")
	//logger.Info("this is prefix ")s
}
