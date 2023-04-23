package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	GlobalCache *cache.Cache
)

func init() {
	GlobalCache = cache.New(1*time.Minute, 10*time.Minute)
}
