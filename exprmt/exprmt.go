// A file for experimenting with things

package main

import (
	"rwb/exprmt/logging"
	"time"
)

func main() {
	logger := logging.New(time.DateTime, false)
	logger.Log("info", "starting up service")
	logger.Log("warning", "no tasks found")
	logger.Log("error", "exiting: no work performed")
}
