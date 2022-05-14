package gingenius

import (
	"github.com/gin-gonic/gin"
)

type GeniusServer interface {
	With(f *Features) (err error)
}

type Server struct {
	GeniusServer
	Engine *gin.Engine
}

type Features struct {
	routes gin.RoutesInfo
}

func NewFeatures(routes gin.RoutesInfo) *Features {
	f := &Features{
		routes: routes,
	}

	return f
}

func (s *Server) With(f *Features) (err error) {
	for _, r := range f.routes {
		s.Engine.Handle(r.Method, r.Path, r.HandlerFunc)
	}

	return err
}
