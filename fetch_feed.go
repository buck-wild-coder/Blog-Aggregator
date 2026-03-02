package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	var rssfeed RSSFeed
	err = xml.Unmarshal(data, &rssfeed)
	if err != nil {
		return err
	}
	rssfeed = unescapestring(rssfeed)
	fmt.Println(rssfeed)
	return nil
}

func unescapestring(rssfeed RSSFeed) RSSFeed {
	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Description)
	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)
	for i := range rssfeed.Channel.Item {
		rssfeed.Channel.Item[i].Title = html.UnescapeString(rssfeed.Channel.Item[i].Title)
		rssfeed.Channel.Item[i].Description = html.UnescapeString(rssfeed.Channel.Item[i].Description)
	}
	return rssfeed
}
