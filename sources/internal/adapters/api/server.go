// HRACH_DEV © iMed Cloud Services, Inc.
package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	port = 9999
)

type server struct {
	http.Server
}

func (s *server) Start() error {
	log.Printf("starting API server on port: %d with %d CPU cores ...", port, runtime.NumCPU())
	if s.TLSConfig == nil {
		return s.ListenAndServe()
	}
	return s.ListenAndServeTLS("", "")
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(ctx); err == nil {
		log.Println("API server stopped")
		return nil
	} else {
		return err
	}
}

func NewAPIServer(hdl *gin.Engine) (*server, error) {
	srv := &server{}
	srv.Addr = fmt.Sprintf(":%d", port)
	srv.Handler = hdl
	return srv, nil
}
