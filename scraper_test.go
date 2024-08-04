package main

import (
	"reflect"
	"testing"
)

func TestIOReadDir(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			"IORead Band",
			"./test/",
			[]string{"band", "band2"},
		},
		{
			"IORead Albums Band 1",
			"./test/band",
			[]string{"album_1", "album_2"},
		},
		{
			"IORead Albums  Band 2",
			"./test/band2",
			[]string{"band_2_album_1", "band_2_album_2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := IOReadDir(tt.args)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IOReadDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocalData(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []libraryStruct
	}{
		{
			"LocalAlbums",
			"./test/",
			[]libraryStruct{{"band", []string{"album_1", "album_2"}, []string{}}, {"band2", []string{"band_2_album_1", "band_2_album_2"}, []string{}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetLocalData(tt.args)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocalAlbums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRemoteAlbums(t *testing.T) {
	tests := []struct {
		name string
		band libraryStruct
		want libraryStruct
	}{
		{
			"RemoteAlbums",
			libraryStruct{name: "band name", localAlbums: []string{"Album 1", "Album 2"}, remoteAlbums: []string{}},
			libraryStruct{name: "band name", localAlbums: []string{"Album 1", "Album 2"}, remoteAlbums: []string{"Album 1", "Album 2", "Album 3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRemoteAlbums(tt.band); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRemoteAlbums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRemoteData(t *testing.T) {
	tests := []struct {
		name    string
		library []libraryStruct
		want    []libraryStruct
	}{
		{
			"RemoteData",
			[]libraryStruct{
				{name: "band name", localAlbums: []string{"Album 1", "Album 2"}, remoteAlbums: []string{}},
				{name: "band 2", localAlbums: []string{"Album 1", "Album 2"}, remoteAlbums: []string{}},
			},
			[]libraryStruct{
				{name: "band name", localAlbums: []string{"Album 1", "Album 2"}, remoteAlbums: []string{"Album 1", "Album 2", "Album 3"}},
				{name: "band 2", localAlbums: []string{"Album 1", "Album 2"}, remoteAlbums: []string{"Album 1", "Album 2", "Album 3"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetRemoteData(tt.library)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRemoteData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareAlbums(t *testing.T) {
	tests := []struct {
		name string
		band libraryStruct
		want albumsStruct
	}{
		{
			"compare new albums",
			libraryStruct{name: "band name", localAlbums: []string{"Album 1", "Album 2"}, remoteAlbums: []string{"Album 1", "Album 2", "Album 3"}},
			albumsStruct{"band name", []string{"Album 3"}},
		},
		{
			"compare no new albums",
			libraryStruct{name: "band", localAlbums: []string{"Album 1", "Album 2", "Album 3"}, remoteAlbums: []string{"Album 1", "Album 2", "Album 3"}},
			albumsStruct{"band", []string{"No new albums"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareAlbums(tt.band); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareAlbums() = %v, want %v", got, tt.want)
			}
		})
	}
}
