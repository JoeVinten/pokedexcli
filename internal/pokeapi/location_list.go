package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (PokemonLocations, error) {

	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	val, cacheHit := c.cache.Get(url)

	if cacheHit {
		// cache hit
		cachedLocations := PokemonLocations{}

		err := json.Unmarshal(val, &cachedLocations)

		if err != nil {
			return PokemonLocations{}, err
		}
		return cachedLocations, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return PokemonLocations{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return PokemonLocations{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return PokemonLocations{}, err

	}

	if !cacheHit {
		c.cache.Add(url, body)
	}

	locations := PokemonLocations{}

	err = json.Unmarshal(body, &locations)

	if err != nil {
		return PokemonLocations{}, err
	}

	return locations, nil
}
