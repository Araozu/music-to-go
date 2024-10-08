package utils

import "time"

// The result of a Login
type AuthSuccess struct {
	Id            string `json:"id"`
	IsAdmin       bool   `json:"isAdmin"`
	Name          string `json:"name"`
	SubsonicSalt  string `json:"subsonicSalt"`
	SubsonicToken string `json:"subsonicToken"`
	Token         string `json:"token"`
	Username      string `json:"username"`
}

// An album as defined by the server
type Album struct {
	PlayCount     int       `json:"playCount"`
	PlayDate      time.Time `json:"playDate"`
	Rating        int       `json:"rating"`
	Starred       bool      `json:"starred"`
	StarredAt     time.Time `json:"starredAt"`
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	EmbedArtPath  string    `json:"embedArtPath"`
	ArtistID      string    `json:"artistId"`
	Artist        string    `json:"artist"`
	AlbumArtistID string    `json:"albumArtistId"`
	AlbumArtist   string    `json:"albumArtist"`
	AllArtistIds  string    `json:"allArtistIds"`
	MaxYear       int       `json:"maxYear"`
	MinYear       int       `json:"minYear"`
	Compilation   bool      `json:"compilation"`
	SongCount     int       `json:"songCount"`
	Duration      float64   `json:"duration"`
	Size          int       `json:"size"`
	Genre         string    `json:"genre"`
	Genres        []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	FullText              string    `json:"fullText"`
	OrderAlbumName        string    `json:"orderAlbumName"`
	OrderAlbumArtistName  string    `json:"orderAlbumArtistName"`
	ImageFiles            string    `json:"imageFiles"`
	Paths                 string    `json:"paths"`
	SmallImageURL         string    `json:"smallImageUrl"`
	MediumImageURL        string    `json:"mediumImageUrl"`
	LargeImageURL         string    `json:"largeImageUrl"`
	ExternalURL           string    `json:"externalUrl"`
	ExternalInfoUpdatedAt time.Time `json:"externalInfoUpdatedAt"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type NavError struct {
	Error string `json:"error"`
}

type Song struct {
	PlayCount        int         `json:"playCount"`
	PlayDate         time.Time   `json:"playDate"`
	Rating           int         `json:"rating"`
	Starred          bool        `json:"starred"`
	StarredAt        interface{} `json:"starredAt"`
	BookmarkPosition int         `json:"bookmarkPosition"`
	ID               string      `json:"id"`
	LibraryID        int         `json:"libraryId"`
	Path             string      `json:"path"`
	Title            string      `json:"title"`
	Album            string      `json:"album"`
	ArtistID         string      `json:"artistId"`
	Artist           string      `json:"artist"`
	AlbumArtistID    string      `json:"albumArtistId"`
	AlbumArtist      string      `json:"albumArtist"`
	AlbumID          string      `json:"albumId"`
	HasCoverArt      bool        `json:"hasCoverArt"`
	TrackNumber      int         `json:"trackNumber"`
	DiscNumber       int         `json:"discNumber"`
	Year             int         `json:"year"`
	OriginalYear     int         `json:"originalYear"`
	ReleaseYear      int         `json:"releaseYear"`
	Size             int         `json:"size"`
	Suffix           string      `json:"suffix"`
	Duration         float64     `json:"duration"`
	BitRate          int         `json:"bitRate"`
	SampleRate       int         `json:"sampleRate"`
	Channels         int         `json:"channels"`
	Genre            string      `json:"genre"`
	Genres           []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	OrderTitle           string    `json:"orderTitle"`
	OrderAlbumName       string    `json:"orderAlbumName"`
	OrderArtistName      string    `json:"orderArtistName"`
	OrderAlbumArtistName string    `json:"orderAlbumArtistName"`
	Compilation          bool      `json:"compilation"`
	Lyrics               string    `json:"lyrics"`
	RgAlbumGain          int       `json:"rgAlbumGain"`
	RgAlbumPeak          int       `json:"rgAlbumPeak"`
	RgTrackGain          int       `json:"rgTrackGain"`
	RgTrackPeak          int       `json:"rgTrackPeak"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
