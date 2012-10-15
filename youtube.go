package main

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

type YoutubeLink struct {
	Title string  `xml:"entry>title"`
	Link  []*Link `xml:"entry>link"`
}

func (l *YoutubeLink) String() string {
	return fmt.Sprintf("%s [ %s ]", l.Link[0].Address, l.Title)
}

type Link struct {
	Address string `xml:"href,attr"`
}

func feelinglucky(body []byte) *YoutubeLink {
	link := new(YoutubeLink)
	xml.Unmarshal(body, &link)
	return link
}

func (song *Song) youtubesearch() string {
	search, _ := url.Parse("https://gdata.youtube.com/feeds/api/videos")
	search.RawQuery = fmt.Sprintf("q=%s&max-results=1&v=2&prettyprint=true", url.QueryEscape(song.Artist+" - "+song.Title))
	return search.String()
}
