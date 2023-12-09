// Erik Petrosyan ©
package system

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func HandleSysCalls(cancelCtx context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	for range sigChan {
		cancelCtx()
		return
	}
}
