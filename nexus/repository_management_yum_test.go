package nexus_test

import (
	"github.com/Kev1nChan/nexus-go-client/models"
	"github.com/stretchr/testify/assert"
)

func (suite *NexusClientSuite) TestRepositoryManagementYUMHosted() {
	name := "yum-hosted-test"
	testRepository := models.YUMHostedRepository{
		Cleanup: &models.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Name: name,
		YUM: &models.YUMAttributes{
			DeployPolicy:  "PERMISSIVE",
			RepodataDepth: 2,
		},
		Online: true,
		Storage: &models.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}

	// Create
	err := suite.client.RepositoryManagement.CreateYUMHosted(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateYUMHosted(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementYUMProxy() {
	name := "test-centos"
	testRepository := models.YUMProxyRepository{
		Name: name,
		Cleanup: &models.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Storage: &models.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: &models.Proxy{
			RemoteURL:      "https://mirrors.aliyun.com/centos/",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: &models.NegativeCache{
			Enabled:    true,
			TimeToLive: 1440,
		},
		HTTPClient: &models.HTTPClient{
			Blocked:   false,
			AutoBlock: true,
		},
	}
	err := suite.client.RepositoryManagement.CreateYUMProxy(testRepository)
	assert.NoError(suite.T(), err)

	err = suite.client.RepositoryManagement.UpdateYUMProxy(name, testRepository)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementYUMGroup() {
	name := "test-centos-group"
	testRepository := models.YUMGroupRepository{
		Name:   name,
		Online: true,
		Storage: &models.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: &models.Group{
			MemberNames: []string{"test-centos"},
		},
	}
	err := suite.client.RepositoryManagement.CreateYUMGroup(testRepository)
	assert.NoError(suite.T(), err)

	err = suite.client.RepositoryManagement.UpdateYUMGroup(name, testRepository)
	assert.NoError(suite.T(), err)
}
