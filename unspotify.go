package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
