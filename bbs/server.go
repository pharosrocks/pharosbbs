package bbs

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func (s *Server) ListenAndServe(addr string) (err error) {

	// cert := GetCertificate()

	// tlsConfig := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// }

	routeGroup := gin.Default()
	routeGroup.GET("/", gin.HandlerFunc(s.telnetHandler))

	wsServer := http.Server{
		//		TLSConfig: tlsConfig,
		Handler: routeGroup,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	err = wsServer.ListenAndServe()
	routeGroup.Run()

	go func(done chan os.Signal) {
		<-done
		wsServer.Close()
	}(done)

	return err
}
