package metadata

type Metadata struct {
	Title    string
	MetaTags []MetaTag
}

type MetaTag struct {
	Name    string `json:name`
	Content string `json:content`
}

// var MetadataPack = []Metadata{
// 	{Title: "The standard Lorem Ipsum passage", Header: "Eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
// 	{Title: "Finibus Bonorum et Malorum", Header: "Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit", Description: "Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"},
// 	{Title: "Carthago delenda est", Header: "Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam", Description: "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."},
// 	{Title: "In vino veritas", Header: "Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur", Description: "Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."},
// }
