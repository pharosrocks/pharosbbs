package bbs

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/appleboy/graceful"
	"github.com/enriquebris/goconcurrentqueue"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ginmills/ginmill"
	"github.com/ginmills/mastodon"
)

type Server struct {
	ginmill.Server
	mastodon.IMastodon

	fifo    *goconcurrentqueue.FIFO
	manager *graceful.Manager
}

func NewServer() *Server {
	s := new(Server)
	s.fifo = goconcurrentqueue.NewFIFO()
	return s
}

func (s *Server) ListenAndServe(addr string) (err error) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// cert := GetCertificate()

	// tlsConfig := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// }

	gin.SetMode(gin.ReleaseMode)
	s.Engine = gin.New()

	// TODO: custom static dir
	s.Engine.Use(
		logger.SetLogger(),
		gin.Recovery(),
	)

	// cheers
	s.With(s.cheers())

	// add bbsd features
	s.With(s.bbsd())

	// add mastodon features
	s.With(s.mastodon())

	wsServer := http.Server{
		//		TLSConfig: tlsConfig,
		Handler: s.Engine,
		Addr:    addr,
	}

	s.manager = graceful.NewManager()

	// wsServer.ListenAndServe
	go func() {
		if err := wsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Str("error", err.Error())
		}
	}()

	// wsServer.Shutdown
	s.manager.AddRunningJob(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				if err := wsServer.Shutdown(ctx); err != nil {
					// TODO: log shutdown failed
					return err
				}
				return nil
			default:
				time.Sleep(1 * time.Second)
			}
		}
	})

	// dispatcher goroutine
	s.manager.AddRunningJob(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				// TODO: s.fifo clean up
				return nil
			default:
				value, _ := s.fifo.DequeueOrWaitForNextElementContext(ctx)

				if value != nil {
					log.Printf("%v", value)
				}
			}
		}
	})

	<-s.manager.Done()

	return err
}
