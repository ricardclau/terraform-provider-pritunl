package client

import (
	"fmt"
)

type LinkServerOrganizationData struct {
	OrganizationId string `json:"id"`
	Server         string `json:"server"`
}

func (c *PritunlClient) LinkServerOrganizationGet(server string, organizationId string) (*LinkServerOrganizationData, error) {

	req := Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s/organization", server),
	}
	var data []LinkServerOrganizationData
	err := c.Do(req, &data)
	if err != nil {
		return nil, err
	}

	for _, serverOrg := range data {
		if serverOrg.OrganizationId == organizationId {
			return &serverOrg, nil
		}
	}

	return nil, fmt.Errorf("Cannot find organizationId: %v in server: %v", organizationId, server)
}

func (c *PritunlClient) LinkServerOrganizationAttach(server string, organizationId string) (*LinkServerOrganizationData, error) {

	req := Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/organization/%s", server, organizationId),
	}

	data := &LinkServerOrganizationData{}

	err := c.Do(req, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *PritunlClient) LinkServerOrganizationDelete(server string, organizationId string) error {

	req := Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/server/%s/organization/%s", server, organizationId),
	}

	return c.Do(req, nil)
}
