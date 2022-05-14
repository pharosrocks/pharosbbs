package mastodon

import (
	"github.com/gin-gonic/gin"
	"github.com/pharosrocks/pharosbbs/gingenius"
)

type IMastodon interface {
	MastodonPublicTimeline(c *gin.Context)
	MastodonHomeTimeline(c *gin.Context)
	MastodonDirectTimeline(c *gin.Context)
}

// mastodon features
func Features(m IMastodon) (features *gingenius.Features) {
	r := gin.New()

	api := r.Group("v1/")
	api.GET("timelines/home", gin.HandlerFunc(m.MastodonPublicTimeline))
	api.GET("timelines/public", gin.HandlerFunc(m.MastodonHomeTimeline))
	api.GET("timelines/direct", gin.HandlerFunc(m.MastodonDirectTimeline))

	features = gingenius.NewFeatures(r.Routes())

	return features
}
