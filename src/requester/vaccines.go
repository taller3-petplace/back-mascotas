package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"petplace/back-mascotas/src/model"
	"petplace/back-mascotas/src/requester/domain"
)

// GetVaccines fetches all the vaccines that were applied to the pet
func (r *Requester) GetVaccines(petID int) ([]model.Vaccine, error) {

	url := fmt.Sprintf("%s/%d", r.baseUrl, petID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	response, err := r.clientHTTP.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error performing getVaccines request: %w", err)
	}
	defer func() {
		if response != nil && response.Body != nil {
			_ = response.Body.Close()
		}
	}()

	if response == nil || err != nil || response.StatusCode != http.StatusOK {
		var errorInfo error
		if err != nil {
			errorInfo = fmt.Errorf(" Error: %w", err)
		} else {
			rawBody := make([]byte, 0)
			_, _ = response.Body.Read(rawBody)
			errorInfo = fmt.Errorf(" Status: %v | Body: %v ", response.Status, string(rawBody))
		}
		return nil, fmt.Errorf("error from treatments service %w", errorInfo)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading vaccines body: %w", err)
	}

	var vaccinesResponse []domain.VaccineResponse
	err = json.Unmarshal(responseBody, &vaccinesResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling vaccines response: %v", err)
	}

	var vaccines []model.Vaccine
	for _, vaccine := range vaccinesResponse {
		vaccines = append(vaccines, vaccine.ToModel())
	}

	return vaccines, nil
}
