package bizlogger

import (
	"github.com/pochard/logrotator"
	"time"
)

func NewTabLogger(pattern string, period time.Duration, sep string) (*TabLogger, error) {
	output, err := logrotator.NewTimeBasedRotator(pattern, period)
	if err != nil {
		return nil, err
	}
	tabLogger := TabLogger{output, sep}
	return &tabLogger, nil
}
