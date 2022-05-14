package bbs

import (
	"github.com/gin-gonic/gin"
	"github.com/pharosrocks/pharosbbs/gingenius"
	"github.com/pharosrocks/pharosbbs/mastodon"
)

func (s *Server) mastodon() (features *gingenius.Features) {
	features = mastodon.Features(s)
	return features
}

func (s *Server) MastodonPublicTimeline(c *gin.Context) {

}

func (s *Server) MastodonHomeTimeline(c *gin.Context) {

}

func (s *Server) MastodonDirectTimeline(c *gin.Context) {

}
