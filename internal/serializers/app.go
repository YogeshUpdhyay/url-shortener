package serializers

import "github.com/YogeshUpdhyay/url-shortener/internal/constants"

type CreateAppRequest struct {
	AppName string `json:"appName"`
}

func (r *CreateAppRequest) Validate() error {
	if r.AppName == constants.Empty {
		return constants.ValidationError
	}
	return nil
}

type CreateAppResponse struct {
	AppName string `json:"appName"`
	ApiKey  string `json:"apiKey"`
	ID      uint   `json:"id"`
}
