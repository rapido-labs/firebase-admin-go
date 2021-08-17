package remoteconfig

import (
	"encoding/json"
	"time"

	"google.golang.org/api/iterator"
)

// TagColor represents a tag color
type TagColor string

// Tag colors
const (
	colorUnspecified TagColor = ""
	Blue                      = "BLUE"
	Brown                     = "BROWN"
	Cyan                      = "CYAN"
	DeepOrange                = "DEEPORANGE"
	Green                     = "GREEN"
	Indigo                    = "INDIGO"
	Lime                      = "LIME"
	Orange                    = "ORANGE"
	Pink                      = "PINK"
	Purple                    = "PURPLE"
	Teal                      = "TEAL"
)

// Version a Remote Config template version.
// Output only, except for the version description.
// Contains metadata about a particular version of the Remote Config template.
// All fields are set at the time the specified Remote Config template is published.
// A version's description field may be specified in PublishTemplate calls
type Version struct {
	Description    string    `json:"description"`
	IsLegacy       bool      `json:"isLegacy"`
	RollbackSource int64    `json:"rollbackSource"`
	UpdateOrigin   string    `json:"updateOrigin"`
	UpdateTime     time.Time `json:"updateTime"`
	UpdateType     string    `json:"updateType"`
	UpdateUser     *User     `json:"updateUser"`
	VersionNumber  int64    `json:"versionNumber,string"`
}

// VersionIterator represents the iterator for looping over versions
type VersionIterator struct{}

// PageInfo represents the information about a Page
func (it *VersionIterator) PageInfo() *iterator.PageInfo {
	// TODO
	return nil
}

// Next will return the next version item in the loop
func (it *VersionIterator) Next() (*Version, error) {
	return nil, nil
}

// ListVersionsResponse is a list of Remote Config template versions
type ListVersionsResponse struct {
	NextPageToken string    `json:"nextPageToken"`
	Versions      []Version `json:"versions"`
}

// ListVersionsOptions to be used as query params in the request to list versions
type ListVersionsOptions struct {
	StartTime        time.Time
	EndTime          time.Time
	EndVersionNumber string
	PageSize         int
	PageToken        string
}

// Condition targets a specific group of users
// A list of these conditions make up part of a Remote Config template
type Condition struct {
	Expression string   `json:"expression"`
	Name       string   `json:"name"`
	TagColor   TagColor `json:"tagColor"`
}

// RemoteConfig represents a Remote Config
type RemoteConfig struct {
	Conditions      []Condition               `json:"conditions"`
	Parameters      map[string]Parameter      `json:"parameters"`
	Version         Version                   `json:"version"`
	ParameterGroups map[string]ParameterGroup `json:"parameterGroups"`
}

// Response to save the API response including ETag
type Response struct {
	*RemoteConfig
	Etag string `json:"etag"`
}

// Parameter .
type Parameter struct {
	ConditionalValues map[string]*ParameterValue `json:"conditionalValues"`
	DefaultValue      *ParameterValue            `json:"defaultValue"`
	Description       string                     `json:"description"`
}

// ParameterValue .
type ParameterValue struct {
	ExplicitValue   string `json:"value"`
	UseInAppDefault bool   `json:"useInAppDefault,omitempty"`
}

// UseInAppDefaultValue returns a parameter value with the in app default as false
func UseInAppDefaultValue() *ParameterValue {
	return &ParameterValue{
		UseInAppDefault: false,
	}
}

// NewExplicitParameterValue will add a new explicit parameter value
func NewExplicitParameterValue(value string) *ParameterValue {
	pm := UseInAppDefaultValue()
	pm.ExplicitValue = value
	return pm
}

// ParameterGroup representing a Remote Config parameter group
// Grouping parameters is only for management purposes and does not affect client-side fetching of parameter values
type ParameterGroup struct {
	Description string                `json:"description"`
	Parameters  map[string]*Parameter `json:"parameters"`
}

// Template .
type Template struct {
	Conditions      []Condition
	ETag            string
	Parameters      map[string]Parameter
	ParameterGroups map[string]ParameterGroup
	Version         Version
}
func(t *RemoteConfig)Mime()string{
	return "application/json"
}
func (t *RemoteConfig)Bytes()([]byte, error){
	return json.Marshal(t)

}

// User represents a remote config user
type User struct {
	Email    string `json:"email"`
	ImageURL string `json:"imageUrl"`
	Name     string `json:"name"`
}
