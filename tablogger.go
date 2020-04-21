package bizlogger

import (
	"github.com/pochard/logrotator"
	"strings"
)

type TabLogger struct {
	output *logrotator.TimeBasedRotator
	sep    string
}

// output like t1,t2,t3\n
func (tlogger *TabLogger) Log(a []string) error {
	switch len(a) {
	case 0:
		_, err := tlogger.output.Write([]byte(""))
		return err
	case 1:
		_, err := tlogger.output.Write([]byte(a[0]))
		return err
	}
	n := len(tlogger.sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	n++

	var b strings.Builder
	b.Grow(n)
	b.WriteString(a[0])
	for _, s := range a[1:] {
		b.WriteString(tlogger.sep)
		b.WriteString(s)
	}

	b.WriteString("\n")

	_, err := tlogger.output.Write([]byte(b.String()))
	return err
}

func (tlogger *TabLogger) Close() error {
	return tlogger.output.Close()
}
