package app

import (
	"fmt"

	"github.com/levilovelock/magistats/scraper"
)

// Start is the main entry point for starting the Application
func Start() {
	// Create and run scraper
	scraper := scraper.Create()
	_, err := scraper.ScrapeTopLeagueDecks()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No errors!")
	}
}
