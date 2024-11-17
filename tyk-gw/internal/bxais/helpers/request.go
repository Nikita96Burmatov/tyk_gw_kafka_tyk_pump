package helpers

import (
	"fmt"
	"github.com/TykTechnologies/tyk/header"
	"net/http"
	"net/url"
	"strings"
)

type Operation string

type UrlStruct struct {
	Version   string
	Recipient string
	Entity    string
	Operation string
	Path      string
}

func validatePath(parts []string) error {
	//if len(parts) < 4 {
	//	return fmt.Errorf("")
	//}

	return nil
}

func ParseURL(path string) ([]string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("url parse error: %v", err)
	}

	pathParts := strings.Split(u.Path, "/")

	if err := validatePath(pathParts); err != nil {
		return nil, fmt.Errorf("path not validated: %v", err)
	}

	return pathParts, nil
}

func Fetch(r *http.Request) (*http.Response, error) {
	//r = r.WithContext(context.Background())
	r.Header.Set(header.ContentType, header.ApplicationJSON)

	resp, err := http.DefaultClient.Do(r)
	//defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("http request error: %v", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close request body: %v", err)
	}

	return resp, nil
}
