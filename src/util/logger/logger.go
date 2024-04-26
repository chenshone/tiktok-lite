package util

import (
	"fmt"
	"io"
	"os"
	"path"
	"tiktok-lite/src/config"
	"tiktok-lite/src/util"
	"time"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var hostname string

func init() {
	hostname, _ = os.Hostname()

	switch config.EnvCfg.Logger.Level {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN", "WARNING":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	}

	filePath := path.Join(util.GetRootPath(), "log", "tiktok-lite", fmt.Sprintf("%s.log", time.Now().Format("2006-01-02:15:04:05")))
	fileDir := path.Dir(filePath)

	if err := os.MkdirAll(fileDir, os.FileMode(0755)); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.AddHook(logTraceHook{})
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	Logger = log.WithFields(log.Fields{
		"Tied":     config.EnvCfg.Logger.Tied,
		"Hostname": hostname,
		"PodIP":    config.EnvCfg.Pod.IP,
	})
}

var Logger *log.Entry

func LogService(name string) *log.Entry {
	return Logger.WithFields(log.Fields{
		"Service": name,
	})
}

type logTraceHook struct{}

func (t logTraceHook) Levels() []log.Level {
	return log.AllLevels
}

func (t logTraceHook) Fire(entry *log.Entry) error {
	ctx := entry.Context
	if ctx == nil {
		return nil
	}

	span := trace.SpanFromContext(ctx)

	sCtx := span.SpanContext()

	if sCtx.HasTraceID() {
		entry.Data["trace_id"] = sCtx.TraceID().String()
	}
	if sCtx.HasSpanID() {
		entry.Data["span_id"] = sCtx.SpanID().String()
	}

	if config.EnvCfg.Logger.WithTranceState == "enable" {
		attrs := make([]attribute.KeyValue, 0)
		logServerityKey := attribute.Key("log.serverity")
		logMessageKey := attribute.Key("log.message")
		attrs = append(attrs, logServerityKey.String(entry.Level.String()))
		attrs = append(attrs, logMessageKey.String(entry.Message))

		for k, v := range entry.Data {
			fields := attribute.Key(fmt.Sprintf("log.fields.%s", k))
			attrs = append(attrs, fields.String(fmt.Sprintf("%v", v)))
		}
		span.AddEvent("log", trace.WithAttributes(attrs...))

		if entry.Level <= log.ErrorLevel {
			span.SetStatus(codes.Error, entry.Message)
		}
	}

	return nil
}
