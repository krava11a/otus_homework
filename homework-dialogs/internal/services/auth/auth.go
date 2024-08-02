package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type res struct {
	Id      string `json:"userId"`
	Message string `json:"message"`
}

type AuthRemoteService struct {
	authPath string
}

func New(authPath string) *AuthRemoteService {
	return &AuthRemoteService{authPath: authPath}
}

func (ars *AuthRemoteService) GetUUIDBy(token, xid string) (id string, err error) {

	response := res{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", ars.authPath, token), nil)
	if err != nil {
		return "", fmt.Errorf("Error ms_dialogs.auth in request ID:%s. Error:%s", xid, err)
	}
	req.Header.Set("X-Request-Id", xid)

	// resp, err := http.Get(fmt.Sprintf("%s/%s", ars.authPath, token))
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error ms_dialogs.auth in request ID:%s. Error:%s", xid, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 500 {
		return "", fmt.Errorf(response.Message)

	}
	id = response.Id
	return

}
