package main

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"
)

type Song struct {
	Artist string `xml:"artist>name"`
	Title  string `xml:"name"`
	link   *YoutubeLink
}

func (song *Song) String() (s string) {
	if song == nil || (*youtube && song.link == nil) {
		return "No results"
	}
	if *youtube {
		return fmt.Sprintf("%s", song.link)
	}
	return fmt.Sprintf("%s - %s", song.Artist, song.Title)
}

func parseSong(body []byte) *Song {
	song := new(Song)
	xml.Unmarshal(body, &song)
	return song
}
func spotifyinfo(link string) string {
	address, _ := url.Parse(link)
	parts := strings.Split(address.Path, "/")

	info, _ := url.Parse("http://ws.spotify.com/lookup/1/")
	info.RawQuery = fmt.Sprintf("uri=spotify:%s:%s", parts[1], parts[2])
	return info.String()
}
