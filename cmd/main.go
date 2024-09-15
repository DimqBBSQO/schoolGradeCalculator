// main package starts copmonents and server
package main

import (
	"context"
	"os"
	"os/signal"
	"schoolGradeCalculator/pkg/logger"
	"syscall"
)

// init initializes the logger by calling ZapLoggerInit function.
func init() {
	logger.ZapLoggerInit()
}

// @title School grade calculator
// @version 1.0
// @description server backend
// @licensename Non-free
func main() {
	defer logger.Sync()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// init context
	ctx, cancel := context.WithCancel(context.Background())

	// cancel context on shutdown
	go func() {
		<-quit
		cancel()
	}()
}
