package main

type Version struct {
	AppVersion string `json:"version"`
	Commit     string `json:"commit"`
	BuildTime  string `json:"buildTime"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Book struct {
	Authors []Author `json:"authors"`
}

type Author struct {
	Key string `json:"key"`
}

type AuthorResponse struct {
	FullName string `json:"fuller_name"`
}

type EndpointAuthorResponse struct {
	AuthorName string `json:"author"`
	Key        string `json:"authorKey"`
}
