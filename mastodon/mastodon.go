package mastodon

import (
	"github.com/gin-gonic/gin"
	"github.com/pharosrocks/pharosbbs/gingenius"
)

type IMastodon interface {
	PublicTimeline(c *gin.Context)
	HomeTimeline(c *gin.Context)
	DirectTimeline(c *gin.Context)
}

type Mastodon struct {
	IMastodon
}

// mastodon features
func Features(m *Mastodon) (features *gingenius.Features) {
	r := gin.New()

	api := r.Group("v1/")
	api.GET("timelines/home", gin.HandlerFunc(m.PublicTimeline))
	api.GET("timelines/public", gin.HandlerFunc(m.HomeTimeline))
	api.GET("timelines/direct", gin.HandlerFunc(m.DirectTimeline))

	features = gingenius.NewFeatures(r.Routes())

	return features
}
