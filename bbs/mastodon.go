package bbs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginmills/ginmill"
	"github.com/ginmills/mastodon"
)

func (s *Server) mastodon() (features *ginmill.Features) {
	features = mastodon.Features(s)
	return features
}

func (s *Server) OAuthAuthorize(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (s *Server) OAuthObtainToken(c *gin.Context) {

}

func (s *Server) OAuthRevokeToken(c *gin.Context) {

}
