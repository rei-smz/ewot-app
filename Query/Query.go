package Query

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type SPARQLConnector struct {
	endpoint string
}

var instance *SPARQLConnector

func GetSPARQLConnector() *SPARQLConnector {
	if instance == nil {
		endpoint := os.Getenv("EWOT_ENDPOINT")
		if endpoint == "" {
			log.Fatal("environment variable EWOT_ENDPOINT is not set")
			return nil
		}
		//newRepo, _ := sparql.NewRepo(endpoint)
		instance = &SPARQLConnector{endpoint: endpoint}
	}
	return instance
}

func (c *SPARQLConnector) SendQuery(query string) (string, error) {
	u, _ := url.Parse(c.endpoint)
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("Accept", "application/sparql-results+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode/100 != 2 {
		return "", fmt.Errorf("sparql GET %s -> %d: %s", u.String(), resp.StatusCode, string(b))
	}
	return string(b), nil
}
