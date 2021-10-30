package io

import (
	"time"

	"github.com/getsentry/sentry-go"
)

type SentryLog struct{}

func InitSentry(dsn string, env string) {
	sentryClient := sentry.ClientOptions{
		Dsn:         dsn,
		Environment: env,
	}
	sentry.Init(sentryClient)
}

func (sl *SentryLog) CaptureError(err error) {
	sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)
}

func (sl *SentryLog) CaptureMessage(msg string) {
	sentry.CaptureMessage(msg)
}
