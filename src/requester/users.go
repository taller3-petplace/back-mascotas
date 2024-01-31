package requester

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"petplace/back-mascotas/src/requester/domain"
)

const (
	baseURL  = "https://api.lnt.digital/users"
	endpoint = "/telegram_id/{telegramID}"
	getUser  = "get_user"
)

// GetUserData fetches information about a user based on the given telegramID
func (r *Requester) GetUserData(telegramID int) (*domain.UserInfo, error) {

	url := fmt.Sprintf("%s/%d", r.baseUrl, telegramID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	response, err := r.clientHTTP.Do(request)
	if err != nil {
		logrus.Errorf("error performing GetUserData: %v", err)
		return nil, fmt.Errorf("error performing GetUserData: %w", err)
	}
	defer func() {
		if response != nil && response.Body != nil {
			_ = response.Body.Close()
		}
	}()

	if response == nil {
		logrus.Error("nil response from users service")
		return nil, fmt.Errorf("error from users service. Status: %v", response.StatusCode)
	}

	if err != nil {
		logrus.Errorf("error from users service. Status: %v. Error: %v", response.StatusCode, err)
		return nil, fmt.Errorf("error from users service. Status: %v. Error: %w", response.StatusCode, err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("error reading users body: %v", err)
		return nil, fmt.Errorf("error reading users body: %w", err)
	}

	var userServiceResponse domain.UserServiceResponse
	err = json.Unmarshal(responseBody, &userServiceResponse)
	if err != nil {
		logrus.Errorf("error unmarshalling user data: %v", err)
		return nil, fmt.Errorf("error unmarshalling user data: %w", err)
	}

	return &userServiceResponse.UserData, nil
}
