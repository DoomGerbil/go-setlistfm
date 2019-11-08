package setlistfm

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	// Default to English, but allow for overrides
	defaultLang = "en"

	// The other option is XML.  Ew, no.
	jsonResponseType = "application/json"
)

var (
	validLanguages = []string{
		"en",	// English (Setlist.fm default)
		"es",	// Spanish
		"fr",	// French
		"de",	// German
		"pt",	// Portugese
		"tr",	// Turkish
		"it",	// Italian
		"pl",	// Polish
	}
)

type setlistFmClient struct {
	httpClient *http.Client
	apiKey     string
	userLang   string
	logger log.logger
}

type RequestMethodNotAllowed struct {
	msg string
}

func (s RequestMethodNotAllowed) Error() string {
	return s.msg
}

// Helper function for executing requests against the Setlist.FM API
func (s *setlistFmClient) doRequest() {

}

func (s *setlistFmClient) createRequest(URL string, requestType string, queryParams []string) (*http.Request, error) {
	if requestType != http.MethodGet {
		return nil, RequestMethodNotAllowed{"Only GET requests are supported by the Setlist.fm API"}
	}

	req, err := http.NewRequest(requestType, URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", jsonResponseType)
	req.Header.Add("Accept-Language", s.userLang)
	req.Header.Add("x-api-key", s.apiKey)

	return req, nil
}

// Creates a Setlist.FM API Client that can be used to query the service
func NewSetlistFmClient(apiKey string, userLang string) *setlistFmClient {
	
	// Allow overriding language, but reset to English if an unsupported language is selected.
	if !isValidLang(userLang) {
		log.WithFields(log.Fields{
			"selectedLang": userLang,
		}).Warningf("Unsupported language requested.  Defaulting to %s.", defaultLang)
		userLang = defaultLang
	}
	
	// TODO set up an http.transport
	tr := http.Transport{}
	return &setlistFmClient{
		httpClient: &http.Client{
			Transport: tr,
		},
		apiKey: apiKey,
		userLang: userLang,
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
