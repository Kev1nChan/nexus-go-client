package nexus_test

import (
	"github.com/Kev1nChan/nexus-go-client/models"
	"github.com/stretchr/testify/assert"
)

func (suite *NexusClientSuite) TestTaskList() {
	tasks, err := suite.client.Tasks.List()
	assert.NoError(suite.T(), err)
	for _, task := range tasks {
		_, err := suite.client.Tasks.ListByType(task.Type)
		assert.NoError(suite.T(), err)

		_, err = suite.client.Tasks.ListById(task.ID)
		assert.NoError(suite.T(), err)
	}
}

func (suite *NexusClientSuite) TestTaskRunAndStop() {
	tasks, err := suite.client.Tasks.ListByType("repository.cleanup")
	assert.NoError(suite.T(), err)

	err = suite.client.Tasks.Run(tasks[0].ID)
	assert.NoError(suite.T(), err)

	err = suite.client.Tasks.Stop(tasks[0].ID)
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestTaskCreate() {
	taskInfo := models.TaskAttributes{
		Action: "coreui_Task",
		Method: "create",
		Data: &models.Data{
			{
				ID:                    "NX.coreui.model.Task-1",
				TypeID:                "blobstore.compact",
				Enabled:               true,
				Name:                  "compact-default",
				NotificationCondition: "FAILURE",
				Schedule:              "manual",
				Properties: &models.Properties{
					BlobstoreName: "default",
				},
				TimeZoneOffset: "+08:00",
			},
		},
		Type: "rpc",
		Tid:  21,
	}
	err := suite.client.Tasks.Create(taskInfo)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestTaskDelete() {
	tasks, err := suite.client.Tasks.ListByType("blobstore.compact")
	assert.NoError(suite.T(), err)
	if len(tasks) != 0 {
		for _, task := range tasks {
			err = suite.client.Tasks.Delete(task.ID)
			assert.NoError(suite.T(), err)
		}
	}
}
