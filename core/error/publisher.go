package apperrors

import "errors"

var Publisher = struct {
	InvalidPublisherID error
	PublisherNotFound  error
}{
	InvalidPublisherID: errors.New("publisher_invalid_id"),
	PublisherNotFound:  errors.New("publisher_not_found"),
}
