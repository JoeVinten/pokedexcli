package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest("GET", url, nil)

	val, cacheHit := c.cache.Get(url)

	if cacheHit {
		cachedPokemon := Pokemon{}

		err = json.Unmarshal(val, &cachedPokemon)

		if err != nil {
			return Pokemon{}, nil
		}

		return cachedPokemon, nil
	}

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("pokemon %s not found", pokemonName)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	pokemon := Pokemon{}

	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		return Pokemon{}, nil
	}

	return pokemon, nil

}
