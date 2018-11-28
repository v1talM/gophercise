package main

import (
	"flag"
	"fmt"
	"gophercise/linkParser"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urlFlags := flag.String("url", "http://iwangle.me", "the url you want to get site-map.")
	flag.Parse()
	/**
	   	1. Get the webpage
	   	2. parse all the links on the page
		3. build proper urls with our links
		4. filter out any links with a diff domain
		5. find all pages (BFS)
		6. print out xml
	 */
	pages := get(*urlFlags)
	for _, page := range pages {
		fmt.Println(page)
	}

}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host: reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := linkParser.Parse(r)
	var ret []string
	for _, l := range links {
		switch  {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base + l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}