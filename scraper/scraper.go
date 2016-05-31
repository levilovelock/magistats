package scraper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"github.com/levilovelock/magistats/magic"
)

// Scraper does scraping
type Scraper struct {
}

// Create creates a new scraper object
func Create() *Scraper {
	// initalise db conn
	return new(Scraper)
}

// ScrapeTopLeagueDecks scrapes latest top results, returning a slice of
// LeagueEvents or an error
func (s *Scraper) ScrapeTopLeagueDecks() ([]*magic.LeagueEvent, error) {
	return scrapeLeagueTopDecks()
}

func scrapeLeagueTopDecks() ([]*magic.LeagueEvent, error) {
	// Find latest events
	resp, err := http.Get("http://magic.wizards.com/section-articles-see-more-ajax?l=en&f=9041&fromDate=&toDate=&event_format=0&sort=DESC&word=&offset=0")
	if err != nil {
		return nil, err
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, err
	}

	type WizardsGetLatest struct {
		Status int
		Data   []string
	}

	wizardsGetLatest := new(WizardsGetLatest)
	jsonErr := json.Unmarshal(body, wizardsGetLatest)
	if jsonErr != nil {
		return nil, jsonErr
	}

	events := []*magic.LeagueEvent{}

	for _, d := range wizardsGetLatest.Data {
		event, eventLinkErr := parseEventFromLink(d)
		if eventLinkErr != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return nil, nil
}

func parseEventFromLink(rawEventLink string) (*magic.LeagueEvent, error) {
	z := html.NewTokenizer(strings.NewReader(rawEventLink))

	var link string
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "a" {
				// Found link
				link = t.Attr[0].Val
			}
		}
	}

	if link == "" {
		return nil, errors.New("Could not find link to a league result")
	}

	return parseEventFromDirectURL("http://magic.wizards.com" + link)
}

func parseEventFromDirectURL(url string) (*magic.LeagueEvent, error) {
	// TODO: Add result parsing here
	return nil, nil
}
