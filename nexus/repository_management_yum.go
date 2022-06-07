package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateYUMHosted create YUM hosted repository
//	endpoint: POST /beta/repositories/yum/hosted
//	parameters:
// 		r: YUMHostedRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateYUMHosted(r YUMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateYUMHosted updates YUM hosted repository
//	endpoint: PUT /beta/repositories/yum/hosted/{repositoryName}
//	parameters:
// 		r: YUMHostedRepository config
// 		repositoryName: Name of the repository to update
//	responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateYUMHosted(repositoryName string, r YUMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateYUMProxy create YUM proxy repository
//	endpoint: POST /beta/repositories/yum/proxy
//	parameters:
// 		r: YUMProxyRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateYUMProxy(r YUMProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateYUMProxy updates YUM proxy repository
//	endpoint: PUT /beta/repositories/yum/proxy/{repositoryName}
//	parameters:
// 		r: YUMProxyRepository config
// 		repositoryName: Name of the repository to update
//	responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateYUMProxy(repositoryName string, r YUMProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateYUMGroup create YUM group repository
//	endpoint: POST /beta/repositories/yum/group
//	parameters:
// 		r: YUMGroupRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateYUMGroup(r YUMGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateYUMGroup updates YUM group repository
//	endpoint: PUT /beta/repositories/yum/group/{repositoryName}
//	parameters:
// 		r: YUMGroupRepository config
// 		repositoryName: Name of the repository to update
//	responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateYUMGroup(repositoryName string, r YUMGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
