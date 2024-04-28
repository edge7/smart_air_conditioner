package cache

import (
	"log"
	"time"

	"github.com/maypok86/otter"
)

var cache otter.Cache[string, string]

func init() {
	var err error
	cache, err = otter.MustBuilder[string, string](100).
		CollectStats().
		Cost(
			func(key string, value string) uint32 {
				return 1
			},
		).
		WithTTL(time.Second * 5).
		Build()
	if err != nil {
		panic(err)
	}
	log.Println("cache init DONE")
}

func Get(key string) (string, bool) {
	value, ok := cache.Get(key)
	if !ok {
		return "", ok
	}
	log.Println("Got from cache: ", value)
	return value, ok
}

func Set(key string, value string) {
	log.Println("Setting in cache: ", value)
	cache.Set(key, value)
}
