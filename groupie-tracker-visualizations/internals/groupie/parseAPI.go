package groupie

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ParseJson() (GeneralApiStruct, error) {
	var emptyStruct GeneralApiStruct
	groupParsed, err1 := ParserGroup()
	locationParsed, err2 := ParserLocations()
	datesParsed, err3 := ParserDates()
	relsParsed, err4 := ParserRelations()
	switch {
	case err1 != nil:
		fmt.Println(err1)
		return emptyStruct, err1
	case err2 != nil:
		fmt.Println(err2)
		return emptyStruct, err2
	case err3 != nil:
		fmt.Println(err3)
		return emptyStruct, err3
	case err4 != nil:
		fmt.Println(err4)
		return emptyStruct, err4
	}
	ParsedJson := GeneralApiStruct{
		GroupsAll:    groupParsed,
		LocationsAll: locationParsed,
		DatesAll:     datesParsed,
		RelationsAll: relsParsed,
	}
	return ParsedJson, nil
}

// this function parses JSON for Artists general info
func ParserGroup() (Groups, error) {
	var groupGeneral Groups

	parseGroup, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return groupGeneral, errors.New("error: failed to request artists data\n")
	}
	defer parseGroup.Body.Close()
	body, err := ioutil.ReadAll(parseGroup.Body) // response body is []byte
	if err != nil {
		return groupGeneral, errors.New("error: failed to read data from parsed JSON\n")
	}

	err = json.Unmarshal(body, &groupGeneral)
	if err != nil {
		return groupGeneral, err
	}

	return groupGeneral, nil
}

// this function parses JSON for Locations general info
func ParserLocations() (TitleLocations, error) {
	var locations TitleLocations

	parseLocs, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return locations, errors.New("error: failed to request artists data\n")
	}
	defer parseLocs.Body.Close()
	body, err := ioutil.ReadAll(parseLocs.Body) // response body is []byte
	if err != nil {
		return locations, errors.New("error: failed to read data from parsed JSON\n")
	}

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locations, err
	}

	return locations, nil
}

// this function parses JSON for Dates general info
func ParserDates() (TitleDates, error) {
	var dates TitleDates

	parseDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return dates, errors.New("error: failed to request artists data\n")
	}
	defer parseDates.Body.Close()
	body, err := ioutil.ReadAll(parseDates.Body) // response body is []byte
	if err != nil {
		return dates, errors.New("error: failed to read data from parsed JSON\n")
	}

	err = json.Unmarshal(body, &dates)
	if err != nil {
		fmt.Println(err)
		return dates, err
	}

	return dates, nil
}

// this function parses JSON for Relations info
func ParserRelations() (TitleRelations, error) {
	var relations TitleRelations

	parseRels, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return relations, errors.New("error: failed to request artists data\n")
	}
	defer parseRels.Body.Close()
	body, err := ioutil.ReadAll(parseRels.Body) // response body is []byte
	if err != nil {
		return relations, errors.New("error: failed to read data from parsed JSON\n")
	}

	err = json.Unmarshal(body, &relations)
	if err != nil {
		fmt.Println(err)
		return relations, err
	}

	/*
		the code below replaces "_" with a space " "
	*/
	for i := range relations.IndexRels {
		for key := range relations.IndexRels[i].DateLocs {
			// save the value of a map element
			tempForValue := relations.IndexRels[i].DateLocs[key]
			// create new key name with required modifications
			newKey := strings.ReplaceAll(key, "_", " ")
			// delete old key-value pair
			delete(relations.IndexRels[i].DateLocs, key)
			// add new key-value pair with new name and old value
			relations.IndexRels[i].DateLocs[newKey] = tempForValue
		}
	}

	return relations, nil
}
