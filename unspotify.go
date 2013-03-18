package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var youtube = flag.Bool("youtube", false, "output youtube link for the song")

func Unspotify(link string) *Song {
	body := get(spotifyinfo(link))
	song := parseSong(body)
	if *youtube {
		body = get(song.youtubesearch())
		song.link = feelinglucky(body)
	}
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

var args []string

func main() {
	flag.Parse()
	args = flag.Args()
	if len(args) == 0 || len(args) > 1 {
		fmt.Println(args)
		return
	}
	song := Unspotify(args[0])
	fmt.Println(song)
}
