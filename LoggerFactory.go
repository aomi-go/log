package log

import "github.com/sirupsen/logrus"

func NewLogger(tag string) *logrus.Entry {
	logger := logrus.New()
	logger.SetLevel(logrus.GetLevel())
	entry := logger.WithField("TAG", tag)
	return entry
}
