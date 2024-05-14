package constants

import (
	"errors"
)

var ValidationError = errors.New("Invalid request payload")
var ErrFetchingClaimsFromToken = errors.New("Error fetching claims from the token.")
