# bizlogger

It is a timebased-rollup logger which can output separated-value log files like CSV(Comma Separated Values). It's goroutine safe.
LoggerManger is a new feature to manage a group of loggers setting in json-format configuration.

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


## TabLogger Example
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



## LoggerManger Example

configuration: log.json
{
    "rootdir": "/data/web_log",
    "loggers": {
        "cookiemapping": {
            "filename": "business/%Y%m%d/cookiemapping-%Y%m%d-%H%M.log",
            "period": 20
        },
        "errorTrace": {
            "filename": "error/%Y%m%d/errorTrace-%Y%m%d-%H%M.log",
            "period": 20
        }
    }
}

```go
package main

import (
	"github.com/pochard/bizlogger"
	"io/ioutil"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	bytes, _ := ioutil.ReadFile("./log.json")
	var loggerManager *bizlogger.LoggerManger
	var logger1, logger2 *bizlogger.TabLogger
	var err error
	loggerManager, err = bizlogger.NewLoggerManger(bytes)
	check(err)
	logger1, err = loggerManager.Get("cookiemapping")
	check(err)
	logger2, err = loggerManager.Get("errorTrace")
	check(err)
	defer loggerManager.Close()

	sl1 := make([]string, 4)
	sl1[0] = "410001"
	sl1[1] = "1583893679"
	sl1[2] = "183.199.195.151"
	sl1[3] = "C03FD5AAB4CA"

	sl2 := make([]string, 3)
	sl2[0] = "410002"
	sl2[1] = "1583893679"
	sl2[2] = "error details"

	logger1.Log(sl1)
	logger2.Log(sl2)

}

```

output:

├── business
│   └── 20200421
│       └── cookiemapping-20200421-1020.log
└── error
    └── 20200421
        └── errorTrace-20200421-1020.log
