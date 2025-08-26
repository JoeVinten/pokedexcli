package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetAreaDetails(areaName string) (AreaDetails, error) {

	url := baseURL + "/location-area/" + areaName

	val, cacheHit := c.cache.Get(url)

	if cacheHit {
		cachedDetails := AreaDetails{}

		err := json.Unmarshal(val, &cachedDetails)

		if err != nil {
			return AreaDetails{}, err
		}

		return cachedDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return AreaDetails{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return AreaDetails{}, nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return AreaDetails{}, nil
	}

	c.cache.Add(url, body)

	details := AreaDetails{}

	err = json.Unmarshal(body, &details)

	if err != nil {
		return AreaDetails{}, err
	}

	return details, nil
}
