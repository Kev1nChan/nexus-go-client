package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RealmsAPI api

type Realms []string

// List realms
//	endpoint: GET /beta/security/realms/active
//	responses:
// 		200: Successful operation
// 		401: Authentication required
//  return:
//		Realms: array list
func (a RealmsAPI) List() (Realms, error) {
	path := fmt.Sprintf("beta/security/realms/active")
	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	var realms Realms
	err = json.Unmarshal(resp, &realms)
	return realms, err
}

// Update realms
//	endpoint: PUT /beta/security/realms/active
//	responses:
// 		200: Successful operation
// 		401: Authentication required
func (a RealmsAPI) Update(realms Realms) error {
	path := fmt.Sprintf("beta/security/realms/active")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(realms)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)

	return err
}
