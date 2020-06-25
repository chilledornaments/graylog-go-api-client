package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateGraylogUser creates a new Graylog user
func (c *Client) CreateGraylogUser(item *UserItem) error {
	var roleArray []string
	var permissionsArray []string

	for _, r := range item.Roles {
		roleArray = append(roleArray, r.(string))
	}

	for _, p := range item.Permissions {
		permissionsArray = append(permissionsArray, p.(string))
	}

	g := &graylogUserItem{
		Username:    item.Username,
		Password:    item.Password,
		Email:       item.Email,
		FullName:    item.FullName,
		Roles:       roleArray,
		Permissions: permissionsArray,
	}

	b, err := json.Marshal(g)

	if err != nil {
		log.Printf("Received error creating %s JSON %s", "CreateGraylogUser", err.Error())
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users", c.apiBaseURL), bytes.NewReader(b))
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-Requested-By", "Terraform-Provider-Graylog")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "CreateGraylogUser", err.Error())
		return err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "CreateGraylogUser", err.Error())
		return err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "CreateGraylogUser", resp.StatusCode)
	log.Printf("%s response body: %s", "CreateGraylogUser", string(rb))

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "CreateGraylogUser", err.Error())
		return err
	}

	switch s := resp.StatusCode; s {
	case 401:
		return ErrGraylogUnauthorized
	case 415:
		return ErrGraylogUnsupportedMedia
	case 400:
		return ErrGraylogBadRequest
	}

	return nil
}

// GetGraylogUser returns a Graylog user's account information
func (c *Client) GetGraylogUser(username string) (*graylogUserItem, error) {

	g := &graylogUserItem{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", c.apiBaseURL, username), nil)
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "GetGraylogUser", err.Error())
		return g, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "GetGraylogUser", err.Error())
		return g, err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "GetGraylogUser", resp.StatusCode)
	log.Printf("%s response body: %s", "GetGraylogUser", string(rb))

	err = json.Unmarshal(rb, g)

	if err != nil {
		log.Printf("%s", err.Error())
		return g, err
	}

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "GetGraylogUser", err.Error())
		return g, err
	}

	switch s := resp.StatusCode; s {
	case 401:
		return g, ErrGraylogUnauthorized
	case 415:
		return g, ErrGraylogUnsupportedMedia
	case 400:
		return g, ErrGraylogBadRequest
	}

	return g, nil

}

// UpdateGraylogUser updates a Graylog user's account settings
func (c *Client) UpdateGraylogUser(item *UserItem) error {

	var roleArray []string
	var permissionsArray []string

	for _, r := range item.Roles {
		roleArray = append(roleArray, r.(string))
	}

	for _, p := range item.Permissions {
		permissionsArray = append(permissionsArray, p.(string))
	}

	g := &graylogUserItem{
		Username:    item.Username,
		Password:    item.Password,
		Email:       item.Email,
		FullName:    item.FullName,
		Roles:       roleArray,
		Permissions: permissionsArray,
	}

	b, err := json.Marshal(g)

	if err != nil {
		log.Printf("Received error creating %s JSON %s", "CreateGraylogUser", err.Error())
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/%s", c.apiBaseURL, item.Username), bytes.NewBuffer(b))
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-Requested-By", "Terraform-Provider-Graylog")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "UpdateGraylogUser", err.Error())
		return err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "UpdateGraylogUser", err.Error())
		return err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "UpdateGraylogUser", resp.StatusCode)
	log.Printf("%s response body: %s", "UpdateGraylogUser", string(rb))

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "UpdateGraylogUser", err.Error())
		return err
	}

	switch s := resp.StatusCode; s {
	case 401:
		return ErrGraylogUnauthorized
	case 415:
		return ErrGraylogUnsupportedMedia
	case 400:
		return ErrGraylogBadRequest
	}

	return nil
}

// DeleteGraylogUser removes a Graylog user
func (c *Client) DeleteGraylogUser(username string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/users/%s", c.apiBaseURL, username), nil)
	req.SetBasicAuth(c.token, "token")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-Requested-By", "Terraform-Provider-Graylog")

	if err != nil {
		log.Printf("Received error creating %s HTTP request %s", "DeleteGraylogUser", err.Error())
		return err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Printf("Received error doing %s HTTP request %s", "DeleteGraylogUser", err.Error())
		return err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s response status code: %d", "DeleteGraylogUser", resp.StatusCode)
	log.Printf("%s response body: %s", "DeleteGraylogUser", string(rb))

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Received error reading %s HTTP response body %s", "DeleteGraylogUser", err.Error())
		return err
	}

	switch s := resp.StatusCode; s {
	case 401:
		return ErrGraylogUnauthorized
	case 415:
		return ErrGraylogUnsupportedMedia
	case 400:
		return ErrGraylogBadRequest
	}

	return nil
}
