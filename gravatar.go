package go_gravatar

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

const (
	gravatarWebsite = "https://gravatar.com"
)

var (
	baseurl url.URL
)

func init() {
	base, _ := url.Parse(gravatarWebsite)
	baseurl = *base
}

// Gravatar
type Gravatar struct {
	Size           int
	DefaultPicture string
}

// hash does the following:
// 1. trim whitespace
// 2. convert everything to lowercase
// 3. calculate the md5sum, which returns a [16]byte
// 5. convert the byte array to a slice [:]
func (g *Gravatar) hash(email []byte) []byte {
	data := md5.Sum(bytes.ToLower(bytes.TrimSpace(email)))
	return data[:]
}

// HashString returns a string-backed md5sum of the email address, using Hash(email []byte)
func (g *Gravatar) HashString(email string) string {
	return hex.EncodeToString(g.hash([]byte(email)))
}

// AvatarUrl provides the avatar url for a given email
func (g *Gravatar) AvatarUrl(email string) string {
	loc := baseurl
	loc.Path = fmt.Sprintf("/avatar/%x", g.hash([]byte(email)))

	var values []string
	if g.DefaultPicture != "" {
		values = append(values, fmt.Sprintf("d=%s", g.DefaultPicture))
	}

	if g.Size != 0 {
		values = append(values, fmt.Sprintf("s=%d", g.Size))
	}

	if len(values) > 0 {
		loc.RawQuery = strings.Join(values, "&")
	}
	return loc.String()
}
