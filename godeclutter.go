package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/purell"
)

// var strFlag = flag.String("long-string", "", "Description")
var preferHttpsFlag = flag.Bool("p", false, "Prefer HTTPS - If there's a https url present, don't print the http for it. (since it will probably just redirect to https)")
var normalizeURLFlag = flag.Bool("c", true, "Clean URLs - Aggressively clean/normalize URLs before outputting them.")
var blacklistExtensionsFlag = flag.Bool("b", true, "Blacklist Extensions - clean some uninteresting extensions.")

var blacklistedExtensions = []string{"css", "scss", "png", "jpg", "jpeg", "img", "svg", "ico", "webp", "webm", "tif", "ttf", "tiff", "otf", "woff", "woff2", "gif", "pdf", "bmp", "eot", "mp3", "mp4", "m4a", "m4p", "avi", "flv", "swf", "eot"} 

func iterInput(c chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c <- scanner.Text()
	}

	close(c)
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func stringInSlice(a string, list []string) (int, bool) {
	for i, b := range list {
		if b == a {
			return i, true
		}
	}
	return 0, false
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func normalizeURL(url string) string {
	normalized := purell.MustNormalizeURLString(url, purell.FlagLowercaseScheme|purell.FlagLowercaseHost|purell.FlagUppercaseEscapes|purell.FlagDecodeUnnecessaryEscapes|purell.FlagEncodeNecessaryEscapes|purell.FlagRemoveDefaultPort|purell.FlagRemoveEmptyQuerySeparator|purell.FlagRemoveDotSegments|purell.FlagRemoveDuplicateSlashes|purell.FlagSortQuery)
	return normalized

}

func main() {

	flag.Parse()

	var c chan string = make(chan string)
	go iterInput(c)

	var processedUrls []string
	processedUrlMap := make(map[string]string)

	for line := range c {

		u, err := url.ParseRequestURI(line)
		if err != nil {
			continue
		}

		// Remove http-on-https false-positives
		if u.Scheme == "http" && u.Port() == "443" {
			continue
		}
		if u.Scheme == "https" && u.Port() == "80" {
			continue
		}

		// Escape redundant port syntax (for default ports)
		if strings.Contains(u.Host, ":") {
			host, port, err := net.SplitHostPort(u.Host)
			if err != nil {
				continue
			}

			if port == "443" && u.Scheme == "https" {
				u.Host = host
			} else if port == "80" && u.Scheme == "http" {
				u.Host = host
			} else if port == "" {
				fmt.Print(host)
				u.Host = host
			} else {
				u.Host = host + ":" + port
			}

		}

		// Prefer https
		if *preferHttpsFlag {
			if u.Scheme == "https" {
				scheme, hostname_found := processedUrlMap[u.Host]
				if hostname_found {
					if scheme == "http" {
						check_u, _ := url.Parse(u.String())
						check_u.Scheme = "http"
						found_index, found := stringInSlice(check_u.String(), processedUrls)
						if found {
							processedUrls = remove(processedUrls, found_index)
						}
						processedUrlMap[u.Host] = u.Scheme
					}
				} else {
					processedUrlMap[u.Host] = u.Scheme
				}

			} else if u.Scheme == "http" {
				_, hostname_found := processedUrlMap[u.Host]
				if hostname_found {
					continue
				}
			}
		}

		if *blacklistExtensionsFlag {
			foundBlacklisted := false
			for _, ext := range blacklistedExtensions {
				if strings.HasSuffix(u.Path, ext) {
					foundBlacklisted = true
					continue
				}
			}
			if foundBlacklisted {
				continue
			}
		}

		if *normalizeURLFlag {
			u_str := normalizeURL(u.String())
			processedUrls = append(processedUrls, u_str)
		} else {
			processedUrls = append(processedUrls, u.String())

		}

	}

	filteredProcessedUrls := removeDuplicateStr(processedUrls)

	for _, url := range filteredProcessedUrls {
		fmt.Println(url)
	}

}
