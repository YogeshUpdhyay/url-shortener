package serializers

import (
	"github.com/YogeshUpdhyay/url-shortener/internal/constants"
	"github.com/YogeshUpdhyay/url-shortener/internal/models"
)

type CreateAppRequest struct {
	AppName string `json:"appName"`
}

func (r *CreateAppRequest) Validate() error {
	if r.AppName == constants.Empty {
		return constants.ErrValidatiingRequest
	}
	return nil
}

type CreateAppResponse struct {
	AppName string `json:"appName"`
	ApiKey  string `json:"apiKey"`
	ID      uint   `json:"id"`
}

type ListAppResponse struct {
	PaginationResponse
	Results *[]models.Credential `json:"results"`
}
