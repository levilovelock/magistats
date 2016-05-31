package magic

// Card represents a Magic Card
type Card struct {
	Set    string
	Name   string
	Number int
}

// Deck represents a Magic Deck
type Deck struct {
	Mainboard []Card
	Sideboard []Card
}

// LeagueEvent represents a Daily Event
type LeagueEvent struct {
	Format string
	// Type should be either competitive or friendly
	Type     string
	Date     string
	Listings []*LeagueEntry
}

// LeagueEntry is a placing someone had in a league
type LeagueEntry struct {
	Deck   Deck
	Player string
	Wins   int
	Losses int
}
