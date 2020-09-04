package tweet

import (
	"github.com/isfk/aio/srv/tweet/proto/tweet"
	"github.com/micro/go-micro/v2/client"
)

// V1 V1
var V1 *v1

type v1 struct {
	Client tweet.TweetService
}

// Init Init
func Init() {
	V1 = &v1{
		Client: tweet.NewTweetService("go.micro.service.tweet", client.DefaultClient),
	}
}
