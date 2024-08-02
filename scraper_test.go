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
			[]string{"band"},
		},
		{
			"IORead Albums",
			"./test/band",
			[]string{"album_1", "album_2"},
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
		want []albumsStruct
	}{
		{
			"LocalAlbums",
			"./test/",
			[]albumsStruct{{"band", []string{"album_1", "album_2"}}},
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
		band string
		want albumsStruct
	}{
		{
			"RemoteAlbums",
			"band name",
			albumsStruct{"band name", []string{"Album 1", "Album 2", "Album 3"}},
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
		name        string
		localAlbums []albumsStruct
		want        []albumsStruct
	}{
		{
			"RemoteData",
			[]albumsStruct{{"band1", []string{"album 1", "album 2"}}, {"band2", []string{"album 1", "album 2", "album 3"}}},
			[]albumsStruct{{"band1", []string{"Album 1", "Album 2", "Album 3"}}, {"band2", []string{"Album 1", "Album 2", "Album 3"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetRemoteData(tt.localAlbums)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRemoteData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareAlbums(t *testing.T) {
	tests := []struct {
		name    string
		albums1 []albumsStruct
		albums2 []albumsStruct
		want    []albumsStruct
	}{
		{
			"compare new albums",
			[]albumsStruct{{"band1", []string{"Album 1", "Album 2"}}, {"band2", []string{"Album 1", "Album 2", "Album 3"}}},
			[]albumsStruct{{"band1", []string{"Album 1", "Album 2", "Album 3"}}, {"band2", []string{"Album 1", "Album 2", "Album 3"}}},
			[]albumsStruct{{"band1", []string{"Album 3"}}, {"band2", []string{"No new albums"}}},
		},
		{
			"compare no new albums",
			[]albumsStruct{{"band1", []string{"Album 1", "Album 2"}}, {"band2", []string{"Album 1", "Album 2", "Album 3"}}},
			[]albumsStruct{{"band1", []string{"Album 1", "Album 2"}}, {"band2", []string{"Album 1", "Album 2", "Album 3"}}},
			[]albumsStruct{{"band1", []string{"No new albums"}}, {"band2", []string{"No new albums"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareAlbums(tt.albums1, tt.albums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareAlbums() = %v, want %v", got, tt.want)
			}
		})
	}
}
