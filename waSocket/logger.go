package waSocket

import (
	"fmt"
	"strings"
	"time"
	"wa/api"

	waLog "go.mau.fi/whatsmeow/util/log"
)

type customLogger struct {
	Mod   string
	Color bool
	Min   int
}

var levelToInt = map[string]int{
	"":      -1,
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
}

func (c *customLogger) outputf(level, msg string, args ...interface{}) {
	if levelToInt[level] < c.Min {
		return
	}

	if strings.Compare(level, "ERROR") == 0 {
		api.Passer.Data <- api.SSEData{
			Event:   "notification",
			Message: fmt.Sprintf("%s [%s %s] %s", time.Now().Format("15:04:05.000"), c.Mod, level, fmt.Sprintf(msg, args...)),
		}
	}
}

func (c *customLogger) Errorf(msg string, args ...interface{}) { c.outputf("ERROR", msg, args...) }
func (c *customLogger) Warnf(msg string, args ...interface{})  { c.outputf("WARN", msg, args...) }
func (c *customLogger) Infof(msg string, args ...interface{})  { c.outputf("INFO", msg, args...) }
func (c *customLogger) Debugf(msg string, args ...interface{}) { c.outputf("DEBUG", msg, args...) }

func (c *customLogger) Sub(mod string) waLog.Logger {
	return &customLogger{Mod: fmt.Sprintf("%s/%s", c.Mod, mod), Color: c.Color, Min: c.Min}
}

func LogText(module string, minLevel string, color bool) waLog.Logger {
	return &customLogger{Mod: module, Color: color, Min: levelToInt[strings.ToUpper(minLevel)]}
}
