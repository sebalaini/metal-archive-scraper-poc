package main

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"
)

const PATH = "/Users/sebastiano/Music/Music/Media.localized/Music/"

type albumsStruct struct {
	name   string
	albums []string
}

// The commented keys are part of the response from https://www.metal-archives.com/ API response but not used in the code.
type remoteAlbumsStruct struct {
	// error                string
	// iTotalRecords        int
	// iTotalDisplayRecords int
	// sEcho                int
	aaData [][]string
}

func main() {
	localData, _ := GetLocalData(PATH)
	remoteData, _ := GetRemoteData(localData)

	result := compareAlbums(localData, remoteData)
	for _, band := range result {
		fmt.Printf("%s: %v\n", band.name, band.albums)
	}
}

func IOReadDir(path string) ([]string, error) {
	var files []string
	fileInfo, err := os.ReadDir(path)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		if file.Name() != ".DS_Store" {
			files = append(files, file.Name())
		}
	}

	return files, err
}

func GetLocalData(path string) ([]albumsStruct, error) {
	localData := []albumsStruct{}

	bands, _ := IOReadDir(path)

	for _, band := range bands {
		bandAlbums, _ := IOReadDir(path + band)

		singleBand := albumsStruct{
			name:   band,
			albums: bandAlbums,
		}

		localData = append(localData, singleBand)
	}

	return localData, nil
}

func getRemoteAlbums(band string) albumsStruct {
	// TODO add GET request, the below one is the UI version, while the one used in resp, err := http.Get... is the API URL.
	// metal-archives blocks requests from bots thanks to CloudFlare, so is not possible to fetch real data, neither by using selenium.

	// https://www.metal-archives.com/search/advanced/searching/albums?bandName=band

	// resp, err := http.Get("https://www.metal-archives.com/search/ajax-advanced/searching/albums/?bandName=Vltimas&genre=
	// &country=&yearCreationFrom=&yearCreationTo=&bandNotes=&status=&themes=&location=&bandLabelName=&sEcho=1&iColumns=3&sColumns=
	// &iDisplayStart=0&iDisplayLength=200&mDataProp_0=0&mDataProp_1=1&mDataProp_2=2&_=1639479306089")

	// resp, err := http.Get("https://www.metal-archives.com/search/ajax-advanced/searching/albums/?bandName=band")

	// params := url.Values{}
	// params.Add("bandName", band)
	// resp, err := http.Get("https://www.metal-archives.com/search/ajax-advanced/searching/albums/?" + params.Encode())
	// if err != nil {
	// 	return resp
	// }
	// defer resp.Body.Close()

	rawData := remoteAlbumsStruct{
		// error:                "",
		// iTotalRecords:        2,
		// iTotalDisplayRecords: 2,
		// sEcho:                1,
		aaData: [][]string{
			{
				"<a href='https://www.metal-archives.com/bands/' title='Band'>Band</a>",
				"<a href='https://www.metal-archives.com/albums/'>Album 1</a> <!-- 16.704208 -->",
				"Full-length",
			},
			{
				"<a href='https://www.metal-archives.com/bands/' title='Band'>Band</a>",
				"<a href='https://www.metal-archives.com/albums/'>Album 2</a> <!-- 16.704208 -->",
				"Full-length",
			},
			{
				"<a href='https://www.metal-archives.com/bands/' title='Band'>Band</a>",
				"<a href='https://www.metal-archives.com/albums/'>Album 3</a> <!-- 16.704208 -->",
				"Full-length",
			},
			{
				"<a href='https://www.metal-archives.com/bands/' title='Band'>Band</a>",
				"<a href='https://www.metal-archives.com/albums/'>Demo</a> <!-- 16.704208 -->",
				"Demo",
			},
		},
	}

	albums := []string{}

	for _, album := range rawData.aaData {
		if slices.Contains(album, "Full-length") {
			// TODO add string manipulation to get the URL text out of the HTML string, this is just a temprary solution for the POC
			albums = append(albums, strings.Split(strings.Split(album[1], ">")[1], "</a")[0])
		}
	}

	return albumsStruct{
		name:   band,
		albums: albums,
	}
}

func GetRemoteData(localAlbums []albumsStruct) ([]albumsStruct, error) {
	remoteData := []albumsStruct{}

	for _, band := range localAlbums {
		remoteAlbums := getRemoteAlbums(band.name)

		remoteData = append(remoteData, remoteAlbums)
	}

	return remoteData, nil
}

func GetNewAlbums(localData []albumsStruct, remoteData []albumsStruct) ([]albumsStruct, error) {
	newAlbums := []albumsStruct{}

	for _, band := range newAlbums {
		remoteAlbums := getRemoteAlbums(band.name)

		remoteData = append(remoteData, remoteAlbums)
	}

	return newAlbums, nil
}

func compareAlbums(albums1, albums2 []albumsStruct) []albumsStruct {
	var result []albumsStruct

	for index, band := range albums2 {
		newAlbums := []string{}

		if reflect.DeepEqual(albums1[index].albums, band.albums) {
			newAlbums = append(newAlbums, "No new albums")
		} else {
			for _, album := range band.albums {
				if !slices.Contains(albums1[index].albums, album) {
					newAlbums = append(newAlbums, album)
				}
			}

		}

		result = append(result, albumsStruct{name: band.name, albums: newAlbums})
	}

	return result
}
