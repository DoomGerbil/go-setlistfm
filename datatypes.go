package setlistfm

// An Artist represents a musical artist
// MBID - unique MusicBrainz ID
// TMID - unique Ticketmaster ID
// SortName - non-display name used for sorting (eg "Beatles, The" instead of "The Beatles")
// Disambiguation - freeform string used to disambiguate multiple artists with the same name
// URL - Setlist.fm resource URL
type Artist struct {
	MBID           string `json:"mbid,omitempty"`
	TMID           int    `json:"tmid,omitempty"`
	Name           string `json:"name,omitempty"`
	SortName       string `json:"sortName,omitempty"`
	Disambiguation string `json:"disambiguation,omitempty"`
	URL            string `json:"url,omitempty"`
}

// Artists is a list of Artist elements
type Artists struct {
	Artist       []Artist `json:"artist,omitempty"`
	Total        int      `json:"total,omitempty"`
	Page         int      `json:"page,omitempty"`
	ItemsPerPage int      `json:"itemsPerPage,omitempty"`
}

// Coords represents a geographical location
type Coords struct {
	Long float32 `json:"long,omitempty"`
	Lat  float32 `json:"lat,omitempty"`
}

// A Country represents...a country - duh
type Country struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// Countries is a list of Country elements
type Countries struct {
	Country      []Country `json:"country,omitempty"`
	Total        int       `json:"total,omitempty"`
	Page         int       `json:"page,omitempty"`
	ItemsPerPage int       `json:"items_per_page,omitempty"`
}

// City represents a city - again, duh
type City struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	StateCode   string  `json:"stateCode,omitempty"`
	State       string  `json:"state,omitempty"`
	Coordinates Coords  `json:"coords,omitempty"`
	Country     Country `json:"country,omitempty"`
}

// Cities is a list of City elements
type Cities struct {
	Cities       []City `json:"cities,omitempty"`
	Total        int    `json:"total,omitempty"`
	Page         int    `json:"page,omitempty"`
	ItemsPerPage int    `json:"items_per_page,omitempty"`
}

// An Error represents an error from the server
type Error struct {
	Code      int    `json:"code,omitempty"`
	Status    string `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// A Set represents a Set, which is a subset of a Setlist.
// All Setlists contain one or more Sets
// See https://www.setlist.fm/guidelines#sets for details.
type Set struct {
	Name   string `json:"name,omitempty"`
	Encore int    `json:"encore,omitempty"`
	Song   []Song `json:"song,omitempty"`
}

// A Setlist represents a concert, basically.
// One or more sets of songs, performed by an artist, in a place, at a time.
// Setlists can be edited multiple times, and each edit is represented by a Setlist with a unique VersionID
type Setlist struct {
	Artist      Artist `json:"artist,omitempty"`
	Venue       Venue  `json:"venue,omitempty"`
	Tour        Tour   `json:"tour,omitempty"`
	Set         []Set  `json:"set,omitempty"`
	Info        string `json:"info,omitempty"`
	URL         string `json:"url,omitempty"`
	ID          string `json:"id,omitempty"`
	VersionID   string `json:"version_id,omitempty"`
	EventDate   string `json:"event_date,omitempty"`
	LastUpdated string `json:"last_updated,omitempty"`
}

// Setlists is a list of Setlist elements
type Setlists struct {
	Setlist      []Setlist `json:"setlist,omitempty"`
	Total        int       `json:"total,omitempty"`
	Page         int       `json:"page,omitempty"`
	ItemsPerPage int       `json:"items_per_page,omitempty"`
}

// A Song represents one item played during a Set
type Song struct {
	Name  string `json:"name,omitempty"`
	With  Artist `json:"with,omitempty"`
	Cover Artist `json:"cover,omitempty"`
	Info  string `json:"info,omitempty"`
	Tape  bool   `json:"tape,omitempty"`
}

// A Tour is a collection of concerts (aka Setlists)
// All Tours have at least one Setlist, but not all Setlists are part of a Tour
type Tour struct {
	Name string `json:"name,omitempty"`
}

// A User represents a user of the setlist.fm website
type User struct {
	UserID   string `json:"user_id,omitempty"`
	FullName string `json:"full_name,omitempty"`
	LastFM   string `json:"last_fm,omitempty"`
	MySpace  string `json:"my_space,omitempty"`
	Twitter  string `json:"twitter,omitempty"`
	Flickr   string `json:"flickr,omitempty"`
	Website  string `json:"website,omitempty"`
	About    string `json:"about,omitempty"`
	URL      string `json:"url,omitempty"`
}

// A Venue is the physical location where a Setlist was performed.
type Venue struct {
	City City   `json:"city,omitempty"`
	URL  string `json:"url,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Venues represents a list of Venue elements
type Venues struct {
	Venue        []Venue `json:"venue,omitempty"`
	Total        int     `json:"total,omitempty"`
	Page         int     `json:"page,omitempty"`
	ItemsPerPage int     `json:"items_per_page,omitempty"`
}
