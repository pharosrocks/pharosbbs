package bbs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginmills/cheers"
	"github.com/ginmills/ginmill"
)

func (s *Server) cheers() (features *ginmill.Features) {
	features = cheers.Features(s)
	return features
}

// response to cheers
func (s *Server) Cheers(c *gin.Context) {
	c.JSON(http.StatusOK, "Cheers")
}
