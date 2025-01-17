package pokeclient

import (
	"net/http"
	"time"

	"github.com/rhafaelc/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache *pokecache.Cache
}

func NewClient() Client {
	cache := pokecache.NewCache(5 * time.Second)	

	return Client{
		httpClient: http.Client{},
		cache: cache,
	}
}
