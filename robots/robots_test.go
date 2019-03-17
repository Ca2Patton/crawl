package robots

import (
	"net/http"
	"testing"

	"github.com/temoto/robotstxt"
)

func TestDisallow(t *testing.T) {
	resp, err := http.Get("https://swiftype.com/robots.txt")
	if err != nil {
		t.Error(err)
	}
	robots, err := robotstxt.FromResponse(resp)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Error parsing robots.txt:%s", err)
	}

	ok := robots.TestAgent("/sso", "kitkatattackBot")
	if ok {
		t.Error("Expected False")
	}
}

func TestAllow(t *testing.T) {
	resp, err := http.Get("https://swiftype.com/robots.txt")
	if err != nil {
		t.Error(err)
	}
	robots, err := robotstxt.FromResponse(resp)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Error parsing robots.txt:%s", err)
	}
	ok := robots.TestAgent("/documentation/", "kitkatattackBot")
	if !ok {
		t.Errorf("Expected true")
	}
}
