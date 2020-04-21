package bizlogger

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"time"
)

type Logger struct {
	Filename string `json:"filename"`
	Period   int    `json:"period"`
}

type Config struct {
	Rootdir string `json:"rootdir"`
	Loggers map[string]Logger
}

type LoggerManger struct {
	loggermap map[string]*TabLogger
}

func NewLoggerManger(conf []byte) (*LoggerManger, error) {
	config := Config{}
	err := json.Unmarshal(conf, &config)
	if err != nil {
		return nil, err
	}

	manager := &LoggerManger{}
	manager.loggermap = make(map[string]*TabLogger)
	rootdir := config.Rootdir
	sep := string(os.PathSeparator)
	for name, logger := range config.Loggers {
		tabLogger, err := NewTabLogger(path.Clean(rootdir+sep+logger.Filename), time.Duration(logger.Period)*time.Minute, ",")
		if err != nil {
			return manager, err
		}
		manager.loggermap[name] = tabLogger
	}
	return manager, nil
}

func (lm *LoggerManger) Get(name string) (*TabLogger, error) {
	val, ok := lm.loggermap[name]
	if !ok {
		return nil, errors.New(name + " logger not found")
	}
	return val, nil
}

func (lm *LoggerManger) Close() {
	for _, v := range lm.loggermap {
		v.Close()
	}
}
