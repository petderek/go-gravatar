package go_gravatar

import (
	"testing"
	"net/url"
)

var (
	g Gravatar
)

func TestStringHash(t *testing.T) {
	cases := map[string]string{

		"derek@example.com": "eb23f498f9b14c0e73fd62708f3d2e97",
		// example from their website
		"MyEmailAddress@example.com ": "0bc83cb571cd1c50ba6f3e8a78ef1346",
	}
	for k, v := range cases {
		t.Run("", func(t *testing.T) {
			out := g.HashString(k)
			if v != out {
				t.Errorf("For input (%s) we expected (%s), but actually got (%s).", k, v, out)
			}
		})
	}
}

func TestAvatarUrl(t *testing.T) {
	input := "derek@example.com"
	expected := "https://gravatar.com/avatar/eb23f498f9b14c0e73fd62708f3d2e97"
	result := g.AvatarUrl(input)
	if result != expected {
		t.Errorf("For input (%s) we expected (%s), but actually got (%s).", input, expected, result)
	}
}

func TestAvatarUrlWithParams(t *testing.T) {
	otherG := Gravatar{
		DefaultPicture: "404",
		Size:           200,
	}
	input := "derek@example.com"
	expected := "https://gravatar.com/avatar/eb23f498f9b14c0e73fd62708f3d2e97?d=404&s=200"
	result := otherG.AvatarUrl(input)
	if result != expected {
		t.Errorf("For input (%s) we expected (%s), but actually got (%s).", input, expected, result)
	}
}

func TestAvatarUrlWithNewBase(t *testing.T) {
	domain, _ := url.Parse("https://example.com")
	otherG := Gravatar {
		BaseDomain: domain,
	}

	input := "derek@example.com"
	expected := "https://example.com/avatar/eb23f498f9b14c0e73fd62708f3d2e97"
	result := otherG.AvatarUrl(input)
	if result != expected {
		t.Errorf("For input (%s) we expected (%s), but actually got (%s).", input, expected, result)
	}
}