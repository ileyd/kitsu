package kitsu

import "time"

type Episode struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Titles    struct {
			EnJp string `json:"en_jp"`
		} `json:"titles"`
		CanonicalTitle string      `json:"canonicalTitle"`
		SeasonNumber   int         `json:"seasonNumber"`
		Number         int         `json:"number"`
		RelativeNumber int         `json:"relativeNumber"`
		Synopsis       string      `json:"synopsis"`
		Airdate        string      `json:"airdate"`
		Length         interface{} `json:"length"`
		Thumbnail      struct {
			Original string `json:"original"`
		} `json:"thumbnail"`
	} `json:"attributes"`
	Relationships struct {
		Media struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"media"`
	} `json:"relationships"`
}

// Link returns the full URL to the episode.
func (ep *Episode) Link() string {
	return "https://kitsu.io/episode/" + ep.ID
}
