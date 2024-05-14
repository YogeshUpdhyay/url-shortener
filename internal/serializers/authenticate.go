package serializers

import "github.com/YogeshUpdhyay/url-shortener/internal/constants"

type AuthenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthenticateRequest) Validate() error {
	if r.Password == constants.Empty || r.Email == constants.Empty {
		return constants.ErrValidatiingRequest
	}
	return nil
}

type AuthenticateResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}
