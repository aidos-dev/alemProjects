package groupie

// General struct for initial JSON parsing
type GeneralApiStruct struct {
	GroupsAll    Groups
	LocationsAll TitleLocations
	DatesAll     TitleDates
	RelationsAll TitleRelations
}

// Struct for parsing a group data to be displated on web-page
type GroupDetails struct {
	Group     Group
	Locations Local
	Dates     IndexDate
	Relations IndexRels
}

// ///////////////////////////////////////////////
// Groups
type Groups []Group

type Group struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

//////////////////////////////////////////////////
// Locations

type TitleLocations struct {
	LocsGeneral []Local `json:"index"`
}

type Local struct {
	Id        int
	Locations []string
	Dates     string
}

// ///////////////////////////////////////////////
// Dates

type TitleDates struct {
	Dates []IndexDate `json:"index"`
}

type IndexDate struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}

// ////////////////////////////////////////////////
// Relations

type TitleRelations struct {
	IndexRels []IndexRels `json:"index"`
}

type IndexRels struct {
	Id       int                 `json:"id"`
	DateLocs map[string][]string `json:"datesLocations"`
}

type DateLocs struct {
	Dates []string
}
