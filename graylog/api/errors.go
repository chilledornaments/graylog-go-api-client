package api

import "errors"

var (
	ErrGraylogUnauthorized        = errors.New("Received 401 (Unauthorized) from Graylog. Check your credentials")
	ErrGraylogUnsupportedMedia    = errors.New("Received 415 (Unsupported media type) from Graylog")
	ErrGraylogBadRequest          = errors.New("Received 400 (Bad request) from Graylog. Ensure you've provided valid data")
	ErrGraylogResourceNonExistent = errors.New("Received 404 (Not found) from Graylog. Ensure you're trying to create/access something that exists")
)
