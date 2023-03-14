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
	Title   string   `json:"title"`
}

type Author struct {
	Key string `json:"key"`
}

type AuthorResponse struct {
	FullName string `json:"fuller_name"`
}

type Work struct {
	Title    string `json:"title"`
	Revision int    `json:"revision"`
	Created  struct {
		Value string `json:"value"`
	}
}

type WorksResponse struct {
	Entries []Work `json:"entries"`
}

type EndpointAuthorResponse struct {
	AuthorName string `json:"author"`
	Key        string `json:"authorKey"`
}

type EndpointWorksResponse struct {
	Name        string `json:"name"`
	Revision    int    `json:"revision"`
	PublishDate string `json:"publishDate"`
}
