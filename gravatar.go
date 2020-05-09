package go_gravatar

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

const (
	gravatarWebsite = "https://gravatar.com"
)

var (
	defaultBaseUrl url.URL
)

func init() {
	base, _ := url.Parse(gravatarWebsite)
	defaultBaseUrl = *base
}

// Gravatar
type Gravatar struct {
	Size           int
	DefaultPicture string
	BaseDomain     *url.URL
	Hash           crypto.Hash
}

// hash does the following:
// 1. trim whitespace
// 2. convert everything to lowercase
// 3. calculate the hash, which returns a byte array
// 5. convert the byte array to a slice [:]
func (g *Gravatar) hash(email []byte) (data []byte) {
	sanitized := bytes.ToLower(bytes.TrimSpace(email))
	switch g.Hash {
	case crypto.SHA256:
		d := sha256.Sum256(sanitized)
		data = d[:]
	case crypto.MD5, 0:
		d := md5.Sum(sanitized)
		data = d[:]
	default:
		panic("hash not supported")
	}
	return
}

// HashString returns a string-backed md5sum of the email address, using Hash(email []byte)
func (g *Gravatar) HashString(email string) string {
	return hex.EncodeToString(g.hash([]byte(email)))
}

// AvatarUrl provides the avatar url for a given email
func (g *Gravatar) AvatarUrl(email string) string {
	loc := defaultBaseUrl
	if g.BaseDomain != nil {
		loc = *g.BaseDomain
	}
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
