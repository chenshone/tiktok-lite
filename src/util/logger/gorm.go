package util

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

var errRecordNotFound = errors.New("record not found")

type GormLogger struct{}

func (g GormLogger) LogMode(_ logger.LogLevel) logger.Interface {
	// 我们不使用这个，因为 Gorm 会根据日志集打印不同的日志。
	// 但是，我们只是打印到 TRACE。
	return g
}

func (g GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	Logger.WithContext(ctx).WithFields(logrus.Fields{
		"component": "gorm",
	}).Warnf(s, i...)
}

func (g GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	Logger.WithContext(ctx).WithFields(logrus.Fields{
		"component": "gorm",
	}).Warnf(s, i...)
}

func (g GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	Logger.WithContext(ctx).WithFields(logrus.Fields{
		"component": "gorm",
	}).Errorf(s, i...)
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	const traceStr = "File: %s, Cost: %v, Rows: %v, SQL: %s"
	elapsed := time.Since(begin)
	sql, rows := fc()
	fields := logrus.Fields{
		"component": "gorm",
	}
	if err != nil && errors.Is(err, errRecordNotFound) {
		fields = logrus.Fields{
			"err": err,
		}
	}

	if rows == -1 {
		Logger.WithContext(ctx).WithFields(fields).Tracef(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
	} else {
		Logger.WithContext(ctx).WithFields(fields).Tracef(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}

func GetGormLogger() *GormLogger {
	return &GormLogger{}
}
