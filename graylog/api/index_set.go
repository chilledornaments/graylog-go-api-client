package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) CreateGraylogIndexset(item *GraylogNewIndexSetInput) (string, error) {
	b, err := json.Marshal(item)

	if err != nil {
		log.Printf("Received error creating %s JSON %s", "CreateGraylogIndexset", err.Error())
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/system/indices/index_sets", c.apiBaseURL), bytes.NewReader(b))
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-Requested-By", "Terraform-Provider-Graylog")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "CreateGraylogIndexset", err.Error())
		return "", err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "CreateGraylogIndexset", err.Error())
		return "", err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "CreateGraylogIndexset", resp.StatusCode)
	log.Printf("%s response body: %s", "CreateGraylogIndexset", string(rb))

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "CreateGraylogIndexset", err.Error())
		return "", err
	}

	switch s := resp.StatusCode; s {
	case 401:
		return "", ErrGraylogUnauthorized
	case 415:
		return "", ErrGraylogUnsupportedMedia
	case 400:
		return "", ErrGraylogBadRequest
	case 404:
		return "", ErrGraylogResourceNonExistent
	}

	id := &GraylogInputCreateResponse{}
	err = json.Unmarshal(rb, &id)

	if err != nil {
		return "", err
	}

	return id.ID, nil
}
