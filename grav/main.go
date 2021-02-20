package main

import (
	"crypto"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/petderek/dflag"
	. "github.com/petderek/go-gravatar"
)

var flags = struct {
	Check  bool   `usage:"queries the service for existence. returns the url only if found. exits with (2) if not found"`
	Save   string `usage:"save the photo saves the content to the filename specified. implies check"`
	Size   int    `usage:"size in pixels"`
	Domain string `usage:"base domain to use instead of gravatar, such as http://cdn.libravatar.org"`
	SHA256 bool   `usage:"use sha256 instead of md5, for libravatar and friends"`
}{}

// gravatar's api returns a silhouette by default. passing this arg forces a 404 if an image isn't found
const respondWith404 = "404"

func main() {
	_ = dflag.Parse(&flags)

	if len(dflag.Args()) <= 0 {
		errlog(exitError, "expected first argument to be an email address.")
	}

	var location *url.URL
	if flags.Domain != "" {
		u, err := url.Parse(flags.Domain)
		if err != nil {
			errlog(exitError, "domain was invalid: ", err.Error())
		}
		location = u
	}

	g := &Gravatar{
		Size:           flags.Size,
		DefaultPicture: respondWith404,
		BaseDomain:     location,
	}

	if flags.SHA256 {
		g.Hash = crypto.SHA256
	}

	result := g.AvatarUrl(dflag.Arg(0))
	if flags.Check || flags.Save != "" {
		res, err := http.Get(result)
		if err != nil {
			errlog(exitError, "unable to compete http request: ", err.Error())
		}
		defer res.Body.Close()

		if res.StatusCode == 404 {
			errlog(exitNotFound, "not found")
		}

		if flags.Save != "" {
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				errlog(exitError, "unable to read response: ", err.Error())
			}
			err = ioutil.WriteFile(flags.Save, data, os.ModePerm)
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
