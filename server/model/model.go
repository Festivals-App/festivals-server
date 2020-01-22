package model

import "database/sql"

type Festival struct {
	ID          int    `json:"festival_id"`
	Version     string `json:"festival_version"`
	Valid       bool   `json:"festival_is_valid" db:"festival_is_valid"`
	Name        string `json:"festival_name" db:"festival_name"`
	Start       int    `json:"festival_start" db:"festival_start"`
	End         int    `json:"festival_end" db:"festival_end"`
	Description string `json:"festival_description" db:"festival_description"`
	Price       string `json:"festival_price" db:"festival_price"`
}

func FestivalsScan(rs *sql.Rows) (Festival, error) {
	var f Festival
	return f, rs.Scan(&f.ID, &f.Version, &f.Valid, &f.Name, &f.Start, &f.End, &f.Description, &f.Price)
}

type Artist struct {
	ID          int    `json:"artist_id"`
	Version     string `json:"artist_version"`
	Name        string `json:"artist_name" db:"artist_name"`
	Description string `json:"artist_description" db:"artist_description"`
}

func ArtistsScan(rs *sql.Rows) (Artist, error) {
	var a Artist
	return a, rs.Scan(&a.ID, &a.Version, &a.Name, &a.Description)
}

type Location struct {
	ID          int    `json:"location_id"`
	Version     string `json:"location_version"`
	Name        string `json:"location_name" db:"location_name"`
	Description string `json:"location_description" db:"location_description"`
	Accessible  bool   `json:"location_accessible" db:"location_accessible"`
	Openair     bool   `json:"location_openair" db:"location_openair"`
}

func LocationsScan(rs *sql.Rows) (Location, error) {
	var l Location
	return l, rs.Scan(&l.ID, &l.Version, &l.Name, &l.Description, &l.Accessible, &l.Openair)
}

type Event struct {
	ID          int    `json:"event_id"`
	Version     string `json:"event_version"`
	Name        string `json:"event_name" db:"event_name"`
	Description string `json:"event_description" db:"event_description"`
	Start       int    `json:"event_start" db:"event_start"`
	End         int    `json:"event_end" db:"event_end"`
}

func EventsScan(rs *sql.Rows) (Event, error) {
	var e Event
	return e, rs.Scan(&e.ID, &e.Version, &e.Name, &e.Description, &e.Start, &e.End)
}

type Image struct {
	ID      int    `json:"image_id"`
	Hash    string `json:"image_hash" db:"image_hash"`
	Comment string `json:"image_comment" db:"image_comment"`
	Ref     string `json:"image_ref" db:"image_ref"`
}

func ImagesScan(rs *sql.Rows) (Image, error) {
	var i Image
	return i, rs.Scan(&i.ID, &i.Hash, &i.Comment, &i.Ref)
}

type Link struct {
	ID      int    `json:"link_id"`
	Version string `json:"link_version"`
	URL     string `json:"link_url" db:"link_url"`
	Service int    `json:"link_service" db:"link_service"`
}

func LinksScan(rs *sql.Rows) (Link, error) {
	var l Link
	return l, rs.Scan(&l.ID, &l.Version, &l.URL, &l.Service)
}

type Place struct {
	ID          int     `json:"place_id"`
	Version     string  `json:"place_version"`
	Street      string  `json:"place_street" db:"place_street"`
	ZIP         string  `json:"place_zip" db:"place_zip"`
	Town        string  `json:"place_town" db:"place_town"`
	Addition    string  `json:"place_street_addition" db:"place_street_addition"`
	Country     string  `json:"place_country" db:"place_country"`
	Latitude    float32 `json:"place_lat" db:"place_lat"`
	Longitude   float32 `json:"place_lon" db:"place_lon"`
	Description string  `json:"place_description" db:"place_description"`
}

func PlacesScan(rs *sql.Rows) (Place, error) {
	var p Place
	return p, rs.Scan(&p.ID, &p.Version, &p.Street, &p.ZIP, &p.Town, &p.Addition, &p.Country, &p.Latitude, &p.Longitude, &p.Description)
}

type Tag struct {
	ID   int    `json:"tag_id"`
	Name string `json:"tag_name" db:"tag_name"`
}

func TagsScan(rs *sql.Rows) (Tag, error) {
	var t Tag
	return t, rs.Scan(&t.ID, &t.Name)
}
