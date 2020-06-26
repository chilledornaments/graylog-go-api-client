package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// org.graylog2.inputs.syslog.udp.SyslogUDPInput
func (c *Client) CreateGraylogSyslogUDPGlobalInput(item *GraylogSyslogUDPGlobalInput) (string, error) {
	b, err := json.Marshal(item)

	if err != nil {
		log.Printf("Received error creating %s JSON %s", "CreateGraylogSyslogUDPGlobalInput", err.Error())
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/system/inputs", c.apiBaseURL), bytes.NewReader(b))
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-Requested-By", "Terraform-Provider-Graylog")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "CreateGraylogSyslogUDPGlobalInput", err.Error())
		return "", err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "CreateGraylogSyslogUDPGlobalInput", err.Error())
		return "", err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "CreateGraylogSyslogUDPGlobalInput", resp.StatusCode)
	log.Printf("%s response body: %s", "CreateGraylogSyslogUDPGlobalInput", string(rb))

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "CreateGraylogSyslogUDPGlobalInput", err.Error())
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

func (c *Client) DeleteGraylogSyslogUDPGlobalInput(id string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/system/inputs/%s", c.apiBaseURL, id), nil)
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-Requested-By", "Terraform-Provider-Graylog")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "DeleteGraylogSyslogUDPGlobalInput", err.Error())
		return err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "DeleteGraylogSyslogUDPGlobalInput", err.Error())
		return err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "DeleteGraylogSyslogUDPGlobalInput", resp.StatusCode)
	log.Printf("%s response body: %s", "DeleteGraylogSyslogUDPGlobalInput", string(rb))

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "DeleteGraylogSyslogUDPGlobalInput", err.Error())
		return err
	}

	switch s := resp.StatusCode; s {
	case 401:
		return ErrGraylogUnauthorized
	case 415:
		return ErrGraylogUnsupportedMedia
	case 400:
		return ErrGraylogBadRequest
	case 404:
		return ErrGraylogResourceNonExistent
	}

	return nil
}

func (c *Client) UpdateGraylogSyslogUDPGlobalInput(item *GraylogSyslogUDPGlobalInput, id string) error {

	b, err := json.Marshal(item)

	if err != nil {
		log.Printf("Received error creating %s JSON %s", "UpdateGraylogSyslogUDPGlobalInput", err.Error())
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/system/inputs/%s", c.apiBaseURL, id), bytes.NewReader(b))
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-Requested-By", "Terraform-Provider-Graylog")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "UpdateGraylogSyslogUDPGlobalInput", err.Error())
		return err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "UpdateGraylogSyslogUDPGlobalInput", err.Error())
		return err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "UpdateGraylogSyslogUDPGlobalInput", resp.StatusCode)
	log.Printf("%s response body: %s", "UpdateGraylogSyslogUDPGlobalInput", string(rb))

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "UpdateGraylogSyslogUDPGlobalInput", err.Error())
		return err
	}

	switch s := resp.StatusCode; s {
	case 401:
		return ErrGraylogUnauthorized
	case 415:
		return ErrGraylogUnsupportedMedia
	case 400:
		return ErrGraylogBadRequest
	case 404:
		return ErrGraylogResourceNonExistent
	}

	return nil
}

func (c *Client) GetGraylogSyslogUDPGlobalInput(id string) (*GraylogSyslogUDPGlobalInputGetResponse, error) {
	item := &GraylogSyslogUDPGlobalInputGetResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/system/inputs/%s", c.apiBaseURL, id), nil)
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "CreateGraylogSyslogUDPGlobalInput", err.Error())
		return item, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "CreateGraylogSyslogUDPGlobalInput", err.Error())
		return item, err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "CreateGraylogSyslogUDPGlobalInput", resp.StatusCode)
	log.Printf("%s response body: %s", "CreateGraylogSyslogUDPGlobalInput", string(rb))

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "CreateGraylogSyslogUDPGlobalInput", err.Error())
		return item, err
	}

	switch s := resp.StatusCode; s {
	case 401:
		return item, ErrGraylogUnauthorized
	case 415:
		return item, ErrGraylogUnsupportedMedia
	case 400:
		return item, ErrGraylogBadRequest
	case 404:
		return item, ErrGraylogResourceNonExistent
	}

	err = json.Unmarshal(rb, &item)

	if err != nil {
		return item, err
	}

	return item, nil
}
