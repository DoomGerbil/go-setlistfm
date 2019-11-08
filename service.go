package setlistfm

import (
	"net/http"
)

const (
	defaultLang = "en"
)

var (
	validLanguages = []string{
		"en",
		"es",
		"fr",
		"de",
		"pt",
		"tr",
		"it",
		"pl",
	}
)

type setlistFmClient struct {
	httpClient *http.Client
	apiKey     string
	userLang   string
	acceptType string
}

// Helper function for executing requests against the Setlist.FM API
func (s *setlistFmClient) doRequest() {

}

func (s *setlistFmClient) createRequest(URL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
}

// Creates a Setlist.FM API Client that can be used to query the service
func NewSetlistFmClient(apiKey string, userLang string) *setlistFmClient {
	
	if !isValidLang(userLang) {
		userLang = defaultLang
	}
	
	return &setlistFmClient{
		httpClient: &http.Client{},
		apiKey: apiKey,
		userLang: userLang,
		acceptType: "application/json", // The other option is XML.  Ew.
	}
}

func isValidLang(lang string) bool {
    for _, validLang := range validLanguages {
        if validLang == lang {
            return true
        }
    }
    return false
}
