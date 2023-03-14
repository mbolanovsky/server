package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	version, commit, buildTime string
	// don't edit
	appVersionResponse []byte
)

func sendError(w http.ResponseWriter, httpCode int, err error) {
	buf, locErr := json.Marshal(ErrorResponse{Error: err.Error()})
	if locErr != nil {
		log.Println("origin error", err)
		log.Println("json marshall error", locErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpCode)
	w.Write(buf)
}

func enpGetAPPVersion(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		sendError(w, http.StatusBadRequest, errors.New("unsupported method. GET method is supported only"))
		return
	}

	var err error
	_, err = w.Write(appVersionResponse)
	if err != nil {
		log.Println(err)
	}
}

func enpGetAuthors(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		sendError(w, http.StatusBadRequest, errors.New("unsupported method. GET method is supported only"))
		return
	}

	bookISBN := req.URL.Query().Get("book")
	if bookISBN == "" {
		sendError(w, http.StatusBadRequest, errors.New("missing URL param book which has to contain isbn of book"))
		return
	}

	authors, err := getAuthorsKey(bookISBN)
	if err != nil {
		sendError(w, http.StatusInternalServerError, errors.New("cannot find author of the book"))
		return
	}

	result := make([]EndpointAuthorResponse, len(authors))
	for index := range authors {
		result[index].AuthorName, err = getAuthorsName(authors[index].Key)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err)
			return
		}
		result[index].Key = strings.TrimPrefix(authors[index].Key, "/authors/")
	}

	buf, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		sendError(w, http.StatusInternalServerError, errors.New("cannot create error"))
	}

	_, err = w.Write(buf)
	if err != nil {
		log.Println(err)
	}
}


func enpGetWorks(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		sendError(w, http.StatusBadRequest, errors.New("unsupported method. GET method is supported only"))
		return
	}

	authorId := req.URL.Query().Get("author")
	if authorId == "" {
		sendError(w, http.StatusBadRequest, errors.New("missing URL param book which has to contain author's ID"))
		return
	}

	works, err := getAuthorsKey(authorId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, errors.New("cannot find this author"))
		return
	}

	result := make([]EndpointAuthorResponse, len(authors))
	for index := range authors {
		result[index].AuthorName, err = getAuthorsName(authors[index].Key)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err)
			return
		}
		result[index].Key = strings.TrimPrefix(authors[index].Key, "/authors/")
	}

	buf, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		sendError(w, http.StatusInternalServerError, errors.New("cannot create error"))
	}

	_, err = w.Write(buf)
	if err != nil {
		log.Println(err)
	}
}

func getAuthorsKey(isbn string) ([]Author, error) {
	resp, err := http.Get("https://openlibrary.org/isbn/" + isbn + ".json")
	if err != nil {
		return nil, err
	}

	defer close(resp.Body)

	decoder := json.NewDecoder(resp.Body)

	book := Book{}
	err = decoder.Decode(&book)
	if err != nil {
		return nil, err
	}

	return book.Authors, nil
}

func getAuthorsName(key string) (string, error) {
	resp, err := http.Get("https://openlibrary.org" + key + ".json")
	if err != nil {
		return "", nil
	}

	defer close(resp.Body)

	decoder := json.NewDecoder(resp.Body)

	author := AuthorResponse{}
	err = decoder.Decode(&author)
	return author.FullName, err
}

func close(reader io.ReadCloser) {
	err := reader.Close()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	var err error
	appVersionResponse, err = json.Marshal(&Version{AppVersion: version, BuildTime: buildTime, Commit: commit})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/v1/version", enpGetAPPVersion)
	http.HandleFunc("/api/v1/authors", enpGetAuthors)
	http.HandleFunc("/api/v1/works", enpGetWorks)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
