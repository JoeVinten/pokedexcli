package pokeapi

import (
	"net/http"
	"time"

	"github.com/JoeVinten/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheDuration time.Duration) Client {

	newCache := pokecache.NewCache(cacheDuration)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: newCache,
	}
}
