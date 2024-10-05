package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct{
		Channel struct{
			Title string				`xml:"title"`
			Link string					`xml:"link"`
			Description string	`xml:"description"`
			Item []RSSItem			`xml:"item"`
		} `xml:"channel"`
}

type RSSItem struct{
	Title string				`xml:"title"`
	Link string					`xml:"link"`
	Description string	`xml:"description"`
	PubDate string			`xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 399 {
		return nil, fmt.Errorf("non-OK status code: %v", res.StatusCode)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}

	rssFeed = decodeRSSHTMLEntities(rssFeed)

	return &rssFeed, nil
}

func decodeRSSHTMLEntities(rssFeed RSSFeed) RSSFeed {
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}
	
	return rssFeed
}