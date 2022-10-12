package rule

import "github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/metadata"

type Rule struct {
	Uri      string
	Metadata metadata.Metadata
}

// Dummy data for testing.
var RulePack = []Rule{
	{
		Uri: "https://superbet.com/pariuri-sportive/{sport}",
	},
	{
		Uri: "https://superbet.com/pariuri-sportive/fotbal",
	},
	{
		Uri: "https://superbet.com/pariuri-sportive/{sport}/{country}",
	},
	{
		Uri: "https://superbet.com/pariuri-sportive/fotbal/romania",
	},
	{
		Uri: "https://superbet.com/pariuri-sportive/basket/spain",
	},
	{
		Uri: "https://superbet.com/rezultate/sport/{sport}/{country}/{date}?e={matchId}",
	},
	{
		Uri: "https://superbet.ro/pariuri-sportive/rugby/south-africa",
	},
	{
		Uri: "https://superbet.ro/rezultate/sport/tenis/itf-femei",
	},
	{
		Uri: "https://superbet.ro/rezultate/sport/fotbal/croatia/2022-10-13",
	},
}
