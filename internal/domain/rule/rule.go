package rule

import "github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/metadata"

type Rule struct {
	Path     string // contains the hole Path.
	Metadata metadata.Metadata
}

type Segment struct {
	Rule     Rule
	Children []Segment
}

// Dummy data for testing.
var RulePack = []Rule{
	{
		Path: "https://superbet.com/pariuri-sportive/{sport}",
	},
	{
		Path: "https://superbet.com/pariuri-sportive/fotbal",
	},
	{
		Path: "https://superbet.com/pariuri-sportive/{sport}/{country}",
	},
	{
		Path: "https://superbet.com/pariuri-sportive/fotbal/romania",
	},
	{
		Path: "https://superbet.com/pariuri-sportive/basket/spain",
	},
	{
		Path: "https://superbet.com/rezultate/sport/{sport}/{country}/{date}?e={matchId}",
	},
	{
		Path: "https://superbet.ro/pariuri-sportive/rugby/south-africa",
	},
	{
		Path: "https://superbet.ro/rezultate/sport/tenis/itf-femei",
	},
	{
		Path: "https://superbet.ro/rezultate/sport/fotbal/croatia/2022-10-13",
	},
}

// Dummy data for testing.
var SegmentPack = []Segment{
	{
		Rule: Rule{
			Path: "https://superbet.com",
		},
		Children: []Segment{
			{
				Rule: Rule{
					Path: "/pariuri-sportive",
				},
				Children: []Segment{},
			},
			{
				Rule: Rule{
					Path: "/rezultate",
				},
				Children: []Segment{},
			},
			{
				Rule: Rule{
					Path: "/*",
				},
				Children: []Segment{},
			},
		},
	},
	{
		Rule: Rule{
			Path: "https://superbet.ro",
		},
		Children: []Segment{
			{
				Rule: Rule{
					Path: "/pariuri-sportive",
				},
				Children: []Segment{},
			},
			{
				Rule: Rule{
					Path: "/rezultate",
				},
				Children: []Segment{},
			},
			{
				Rule: Rule{
					Path: "/*",
				},
				Children: []Segment{},
			},
		},
	},
	{
		Rule: Rule{
			Path: "https://superbet.pl",
		},
		Children: []Segment{
			{
				Rule: Rule{
					Path: "/pariuri-sportive",
				},
				Children: []Segment{},
			},
			{
				Rule: Rule{
					Path: "/rezultate",
				},
				Children: []Segment{},
			},
			{
				Rule: Rule{
					Path: "/*",
				},
				Children: []Segment{},
			},
		},
	},
}
