package search

import (
	"encoding/json"
	"os"
)

// Feed struct
type Feed struct {
	Name string `json:"site"`
	URL  string `json:"link"`
	Type string `json:"type"`
}

const file = "data/data.json"

func GetFeeds() ([]*Feed, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var feeds []*Feed
	err = json.NewDecoder(f).Decode(&feeds)
	return feeds, err
}

// TestFeed is Test
// func TestFeed() {
// 	fmt.Println(getFeeds())
// }
