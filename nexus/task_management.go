package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kaizendorks/nexus-go-client/models"
	"net/http"
)

type TaskAPI api

// Create task
//	endpoint: POST extdirect
//	responses:
// 		200: successful operation returns task slice and null error
func (a TaskAPI) Create(t models.TaskAttributes) error {
	path := fmt.Sprintf("extdirect")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(t)
	resp, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	if err != nil {
		return err
	}

	result := make(map[string]interface{})
	json.Unmarshal(resp, &result)
	if result["result"].(map[string]interface{})["success"] == false {
		errInfo := result["result"].(map[string]interface{})["errors"].(map[string]interface{})
		for key, value := range errInfo {
			return fmt.Errorf("%s: %s", key, value)
		}
	}
	return nil
}

// List all task
//	endpoint: GET /v1/tasks
//	responses:
// 		200: successful operation returns task slice and null error
func (a TaskAPI) List() ([]models.Task, error) {
	tasks := models.TaskItems{}
	var items []models.Task

	path := fmt.Sprintf("v1/tasks")

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return items, err
	}
	err = json.Unmarshal(resp, &tasks)
	items = tasks.Items
	return items, err
}

// ListByType task by type
//	endpoint: GET /v1/tasks?type={type}
//	responses:
// 		200: successful operation returns task slice and null error
func (a TaskAPI) ListByType(taskType string) ([]models.Task, error) {
	tasks := models.TaskItems{}
	var items []models.Task

	path := fmt.Sprintf("v1/tasks?type=%s", taskType)

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return items, err
	}
	err = json.Unmarshal(resp, &tasks)
	items = tasks.Items
	return items, err
}

// ListById task by id
//	endpoint: GET /v1/tasks/{id}
//	responses:
// 		200: successful operation returns task slice and null error
//		404: Task not found
func (a TaskAPI) ListById(id string) (models.Task, error) {
	task := models.Task{}

	path := fmt.Sprintf("v1/tasks/%s", id)

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(resp, &task)
	return task, err
}

// Run task run by id
//	endpoint: POST /v1/tasks/{id}/run
//	responses:
// 		204: Task was run
//		404: Task not found
//		405: Task is disabled
func (a TaskAPI) Run(id string) error {
	path := fmt.Sprintf("v1/tasks/%s/run", id)

	_, err := a.client.sendRequest(http.MethodPost, path, nil, nil)
	return err
}

// Stop task stop by id
//	endpoint: POST /v1/tasks/{id}/stop
//	responses:
// 		204: Task was stopped
//		404: Task not found
//		409: Unable to stop task
func (a TaskAPI) Stop(id string) error {
	path := fmt.Sprintf("v1/tasks/%s/stop", id)

	_, err := a.client.sendRequest(http.MethodPost, path, nil, nil)
	return err
}

// Delete task by id
//	endpoint: POST extdirect
//	responses:
// 		200: successful operation returns task slice and null error
//		404: Task not found
func (a TaskAPI) Delete(id string) error {
	path := fmt.Sprintf("extdirect")
	t := models.TaskDeleteAttributes{
		Action: "coreui_Task",
		Method: "remove",
		Data:   []string{id},
		Type:   "rpc",
		Tid:    74,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(t)
	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	if err != nil {
		return err
	}
	return nil
}
