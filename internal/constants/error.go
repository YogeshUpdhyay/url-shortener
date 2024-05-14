package constants

import (
	"errors"
)

var ErrValidatiingRequest = errors.New("invalid request payload")
var ErrFetchingClaimsFromToken = errors.New("error fetching claims from the token")
var ErrPaginationModelCreate = errors.New("error creating pagination model from gin context")
