package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"io/ioutil"
	"net/url"

	. "github.com/petderek/go-gravatar"
)

var (
	check  = flag.Bool("check", false, "queries the service for existence. returns the url only if found. exits with (2) if not found.")
	save   = flag.String("save", "", "save the photo saves the content to the filename specified. implies check")
	size   = flag.Int("size", 0, "the optional size in pixels")
	domain = flag.String("domain", "", "the optional base domain to use instead of gravatar, such as http://cdn.libravatar.org")
)

// gravatar's api returns a silhouette by default. passing this arg forces a 404 if an image isn't found
const respondWith404 = "404"

func main() {
	flag.Parse()

	if flag.NArg() <= 0 {
		errlog(exitError, "expected first argument to be an email address.")
	}

	var location *url.URL
	if *domain != "" {
		u, err := url.Parse(*domain)
		if err != nil {
			errlog(exitError, "domain was invalid: ", err.Error())
		}
		location = u
	}

	g := &Gravatar{
		Size:           *size,
		DefaultPicture: respondWith404,
		BaseDomain:     location,
	}

	result := g.AvatarUrl(flag.Arg(0))
	if *check || *save != "" {
		res, err := http.Get(result)
		if err != nil {
			errlog(exitError, "unable to compete http request: ", err.Error())
		}
		defer res.Body.Close()

		if res.StatusCode == 404 {
			errlog(exitNotFound, "not found")
		}

		if *save != "" {
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				errlog(exitError, "unable to read response: ", err.Error())
			}
			err = ioutil.WriteFile(*save, data, os.ModePerm)
			if err != nil {
				errlog(exitError, "unable to write file: ", err.Error())
			}
		}
	}

	fmt.Println(result)
	os.Exit(0)
}

const (
	_ = iota
	exitError
	exitNotFound
)

func errlog(exitcode int, msg ...string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(exitcode)
}
