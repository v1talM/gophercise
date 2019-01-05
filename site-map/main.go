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
	maxDepth := flag.Int("depth", 3, "the max number of links deep")
	flag.Parse()
	/**
	   	1. Get the webpage
	   	2. parse all the links on the page
		3. build proper urls with our links
		4. filter out any links with a diff domain
		5. find all pages (BFS)
		6. print out xml
	 */
	pages := bfs(*urlFlags, *maxDepth)
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

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for urls, _ := range q {
			if _, ok := seen[urls]; ok {
				continue;
			}
			seen[urls] = struct{}{}
			for _, link := range get(urls) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for urls, _ := range seen {
		ret = append(ret, urls)
	}
	return ret
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