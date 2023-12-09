// Erik Petrosyan ©
package system

import (
	"context"
	"log"
	"sync"
)

type Service interface {
	Stop() error
}

func HandleGracefullExit(ctx context.Context, wg *sync.WaitGroup, services ...Service) {
	<-ctx.Done()
	for _, s := range services {
		if err := s.Stop(); err != nil {
			log.Printf("unable to gracefully stop service (%v)", s)
		}
		wg.Done()
	}
}
