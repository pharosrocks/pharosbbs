package bbs

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/enriquebris/goconcurrentqueue"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Server struct {
	fifo    *goconcurrentqueue.FIFO
	context context.Context
	stop    context.CancelFunc
}

func NewServer() *Server {
	s := new(Server)
	s.fifo = goconcurrentqueue.NewFIFO()
	return s
}

func (s *Server) render() multitemplate.Renderer {
	dir := os.Getenv("WEB_TEMPLATE")

	r := multitemplate.NewRenderer()
	r.AddFromFiles("login", dir+"/login.html")
	return r
}

func (s *Server) ListenAndServe(addr string) (err error) {
	s.context, s.stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer s.stop()

	// cert := GetCertificate()

	// tlsConfig := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// }

	routeGroup := gin.Default()

	// TODO: custom static dir
	routeGroup.Use(static.Serve("/static/", static.LocalFile(os.Getenv("WEB_STATIC"), false)))

	// TODO: custom render dir
	routeGroup.HTMLRender = s.render()

	routeGroup.GET("/", gin.HandlerFunc(s.telnetHandler))
	routeGroup.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login", gin.H{
			"title": "Html5 Template Engine",
		})
	})

	wsServer := http.Server{
		//		TLSConfig: tlsConfig,
		Handler: routeGroup,
		Addr:    addr,
	}

	// wsServer goroutine
	go func() {
		if err := wsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// dispatcher goroutine
	go func() {
		for {
			value, _ := s.fifo.DequeueOrWaitForNextElementContext(s.context)

			if value != nil {
				log.Printf("%v", value)
			}
		}

	}()

	<-s.context.Done()
	log.Printf("shutdown gracefully...")
	s.stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := wsServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	return err
}
