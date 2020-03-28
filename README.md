# bizlogger

It is a timebased-rollup logger which can output separated-value log files like CSV(Comma Separated Values). It's goroutine safe.


Import it in your program as:
```go
      import "github.com/pochard/bizlogger"
```

## API
### pakcage func
	func NewTabLogger(pattern string, period time.Duration, sep string) (*TabLogger, error)

### type TabLogger
	func (tlogger *TabLogger) Log(a []string) error
	func (tlogger *TabLogger) Close() error


## Example
```go
package main

import (
	"fmt"
	"github.com/pochard/bizlogger"
	"time"
)

func main() {
	// rollup every 20 minutes. And the log line is separated by ","
	logger, err := bizlogger.NewTabLogger("/data/web_log/click1-%Y%m%d-%H%M.log", 20*time.Minute, ",")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer logger.Close()

	sl := make([]string, 4)
	sl[0] = "410001"
	sl[1] = "1583893679"
	sl[2] = "183.199.195.151"
	sl[3] = "C03FD5AAB4CA"

	for i := 0; i != 10; i++ {
		err := logger.Log(sl)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
}

```

```
cat /data/web_log/click1-20200328-0940.log
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
410001,1583893679,183.199.195.151,C03FD5AAB4CA
```
