package validators

import (
	"encoding/json"
	"errors"
	"learning/internal/entities"
	"net/http"
)

func ValidateCreateCampaign(r *http.Request) (*entities.CreateCampaignRequest, error) {
	var req entities.CreateCampaignRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return nil, errors.New("invalid JSON payload")
	}

	// Use shared validate instance
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	return &req, nil
}
