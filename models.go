package main

type ErrorResponse struct {
	Error string `json:"error"`
}

type Versions struct {
	Version    string `json:"version"`
	CommitHash string `json:"commitHash"`
	BuildTime  string `json:"buildTime"`
}
