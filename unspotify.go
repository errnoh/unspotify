package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

type Song struct {
	Artist string `xml:"artist>name"`
	Title  string `xml:"name"`
	link   *YoutubeLink
}

func (song *Song) String() string {
	if song == nil || song.link == nil {
		return "No results"
	}
	return fmt.Sprintf("%s", song.link)
}

func (song *Song) youtubesearch() string {
	search, _ := url.Parse("https://gdata.youtube.com/feeds/api/videos")
	search.RawQuery = fmt.Sprintf("q=%s&max-results=1&v=2&orderby=viewCount&prettyprint=true", url.QueryEscape(song.Artist+" - "+song.Title))
	return search.String()
}

func Unspotify(link string) *Song {
	body := get(spotifyinfo(link))
	song := parseSong(body)
	body = get(song.youtubesearch())
	song.link = feelinglucky(body)
	return song
}

func get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic("Failed to GET url")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return body
}

func spotifyinfo(link string) string {
	info, _ := url.Parse("http://ws.spotify.com/lookup/1/")
	parts := strings.Split(link, "/")
	info.RawQuery = fmt.Sprintf("uri=spotify:%s:%s", parts[3], parts[4])
	return info.String()
}

func parseSong(body []byte) *Song {
	song := new(Song)
	xml.Unmarshal(body, &song)
	return song
}

func feelinglucky(body []byte) *YoutubeLink {
	link := new(YoutubeLink)
	xml.Unmarshal(body, &link)
	return link
}

var args []string

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	flag.Parse()
	args = flag.Args()
	if len(args) == 0 || len(args) > 1 {
		fmt.Println(args)
		return
	}
	song := Unspotify(args[0])
	fmt.Println(song)
}
