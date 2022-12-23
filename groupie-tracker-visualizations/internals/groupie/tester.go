package groupie

import "fmt"

// This function tests the API parsing
func TestParser() {
	// Parse groups of artists and print to check if it is parsed correctly
	artist, err := ParserGroup()
	if err != nil {
		fmt.Println("error: couldn't parse artists :(")
		fmt.Println(err)
	}

	fmt.Printf("Artists section:	%v\n", artist[0].Name)

	// Parse locations and print to check if parsed correctly
	local, err := ParserLocations()
	if err != nil {
		fmt.Println("error: couldn't parse locations :(")
		fmt.Println(err)
	}
	fmt.Printf("Locations section:	%v\n", local.LocsGeneral[0].Locations)

	// Parse dates and print to check if parsed correctly
	dates, err := ParserDates()
	if err != nil {
		fmt.Println("error: couldn't parse dates :(")
		fmt.Println(err)
	}
	fmt.Printf("Dates section:		%v\n", dates.Dates[0])

	// Parse relations and print to check if parsed correctly
	relations, err := ParserRelations()
	if err != nil {
		fmt.Println("error: couldn't parse relationsss :(")
		fmt.Println(err)
	}
	fmt.Printf("Relations section:	%v\n", relations.IndexRels[2].Id)
}
