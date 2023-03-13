package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	version, commit, buildTime string
	appVersions                Versions
)

func main() {
	appVersions.Version = version
	appVersions.CommitHash = commit
	appVersions.BuildTime = buildTime

	fmt.Println(appVersions)

	http.HandleFunc("/api/v1/version", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		w.WriteHeader(http.StatusBadRequest)

		resp, err := json.Marshal(ErrorResponse{Error: "unsupported method. support method GET only"})
		if err != nil {
			return
		}
		w.Write(resp)
		return
	}
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
