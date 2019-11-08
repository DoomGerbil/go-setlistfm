package setlistfm

// Artist search results can be sorted one of two ways
type ArtistSortType string
const (
	SortName ArtistSortType = "sortName"
	Relevance ArtistSortType = "relevance"
)

// FindArtistsRequest encapsulates all possible parameters for an artist search
type FindArtistsRequest struct {
	ArtistMBID string
	ArtistName string
	ArtistTMID string
	PageNum int
	SortType ArtistSortType
}

// FindCitiesRequest encapsulates all possible parameters for a city search
type FindCitiesRequest struct {
	Country string
	Name string
	PageNum int
	State string
	StateCode string
}

// FindCountriesRequest is always empty, but is provided for consistency in the Search APIs
type FindCountriesRequest struct {}

// FindSetlistsRequest encapsulates all possible parameters for a setlist search
type FindSetlistsRequest struct {
	ArtistMBID string
	ArtistName string
	ArtistTMID string
	CityID string
	CityName string
	CountryCode string
	Date string
	LastUpdated string
	PageNum int
	State string
	StateCode string
	TourName string
	VenueID string
	VenueName string
	Year string
}

// FindVenuesRequest encapsulates all possible parameters for a venue search
type FindVenuesRequest struct {
	CityID string
	CityName string
	Country string
	Name string
	PageNum int
	State string
	StateCode string
}

// FindArtists searches through available artists 
// Resource: /1.0/search/artists
func (s *setlistFmClient) FindArtists(req FindArtistsRequest) (*Artists, error) {}

// FindCities searches through available cities
// Resource: /1.0/search/cities
func (s *setlistFmClient) FindCities(req FindCitiesRequest) (*Cities, error) {}

// FindCountries lists all available countries
// Resource: /1.0/search/countries
func (s *setlistFmClient) FindCountries(req FindCountriesRequest) (*Countries, error) {}

// FindSetlists searches through available setlists
// Resource: /1.0/search/setlists
func (s *setlistFmClient) FindSetlists(req FindSetlistsRequest) (*Setlists, error) {}

// FindVenues searches through available venues
// Resource: /1.0/search/venues
func (s *setlistFmClient) FindVenues(req FindVenuesRequest) (*Venues, error) {}

// GetArtist fetches details for a specified artist by ID
// Resource: /1.0/artist/{mbid}
func (s *setlistFmClient) GetArtist(artistID string) (*Artist, error) {}

// GetSetlistsForArtist fetches details for a specified artist
// Resource: /1.0/artist/{mbid}/setlists
func (s *setlistFmClient) GetSetlistsForArtist(artistID string, pageNum int) (*Setlists, error) {}

// GetCity fetches details for a specified city
// Resource: /1.0/city/{geoId}
func (s *setlistFmClient) GetCity(cityID string) (*City, error) {}

// GetSetlistByVersion fetches a specified version of a Setlist
// This may not be the final/most recent version.
// Resource: /1.0/setlist/version/{versionId}
func (s *setlistFmClient) GetSetlistByVersion(versionID string) (*Setlist, error) {}

// GetSetlistByID returns the most recent version of a specified setlist
// Resource: /1.0/setlist/{setlistId}
func (s *setlistFmClient) GetSetlistByID(setlistID string) (*Setlist, error) {}

// GetUser fetches details for a specified user
// Resource: /1.0/user/{userId}
func (s *setlistFmClient) GetUser(userID string) (*User, error) {}

// GetSetlistsForUser fetches a list of setlists attended by a specified user
// Resource: /1.0/user/{userId}/attended
func (s *setlistFmClient) GetSetlistsForUser(userID string, pageNum int) (*Setlists, error) {}

// GetEditsForUser fetches a list of setlist edits made by a specified user
// Resource: /1.0/user/{userId}/edited
func (s *setlistFmClient) GetEditsForUser(userID string, pageNum int) (*Setlists, error) {}

// GetVenueByID fetches details for a specified venue
// Resource: /1.0/venue/{venueId}
func (s *setlistFmClient) GetVenueByID(venueID string) (*Venue, error) {}

// GetSetlistsForVenue fetches a list of setlists for a specified venue
// Resource: /1.0/venue/{venueId}/setlists
func (s *setlistFmClient) GetSetlistsForVenue(venueID string, pageNum int) (*Setlists, error) {}
