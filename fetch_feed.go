package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSSFeed struct{
		Channel struct{
			Title string				`xml:"title"`
			Link string					`xml:"link"`
			Description string	`xml:"description"`
			Items []RSSItem			`xml:"item"`
		} `xml:"channel"`
}

type RSSItem struct{
	Title string				`xml:"title"`
	Link string					`xml:"link"`
	Description string	`xml:"description"`
	PubDate string			`xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Add("User-Agent", "gator")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	if res.StatusCode > 399 {
		return &RSSFeed{}, fmt.Errorf("non-OK status code: %v", res.StatusCode)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssFeed := &RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return &RSSFeed{}, err
	}

	return rssFeed, nil
}