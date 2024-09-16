package resources

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type getResourcesResponse struct {
	Resources []Resource `json:"data"`
	Total     int        `json:"total"`
	Page      int        `json:"page"`
	Size      int        `json:"size"`
	Pages     int        `json:"pages"`
}

type Resource struct {
	Name  string         `json:"name"`
	Code  string         `json:"code"`
	Skill string         `json:"skill"`
	Level int            `json:"level"`
	Drops []ResourceDrop `json:"drops"`
}

type ResourceDrop struct {
	Code            string `json:"code"`
	Rate            int    `json:"rate"`
	MinimumQuantity string `json:"min_quantity"`
	MaximumQuantity string `json:"max_quantity"`
}

func GetAllResources() ([]Resource, error) {
	resources := make([]Resource, 0)
	page := 1
	resourceData, err := requestResources(page)
	if err != nil {
		return nil, err
	}

	for page <= resourceData.Pages {
		page++
		resourceData, err := requestResources(page)
		if err != nil {
			return nil, err
		}

		resources = append(resources, resourceData.Resources...)
	}

	return resources, nil
}

func requestResources(page int) (*getResourcesResponse, error) {
	url := "https://api.artifactsmmo.com/resources"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response getResourcesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GetResourceByItemCode(itemCode string) ([]Resource, error) {
	url := fmt.Sprintf("https://api.artifactsmmo.com/resources?drop=%s", itemCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response getResourcesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Resources, nil
}
