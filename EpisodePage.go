package kitsu

import "encoding/json"

type EpisodePage struct {
	Data []Episode `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
}

// GetEpisodePage expects the usual query parameter and returns an EpisodePage object instead of a raw string.
func GetEpisodePage(query string) (*EpisodePage, error) {
	body, requestError := Get(query)

	if requestError != nil {
		return nil, requestError[0]
	}

	page := new(EpisodePage)
	decodeError := json.Unmarshal(body, page)

	return page, decodeError
}
