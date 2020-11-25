package client

import "fmt"

type LinkServerOrganizationPostData struct {
	OrganizationId string `json:"organization_id"`
	Server         string `json:"server"`
}

type LinkServerOrganizationData struct {
	Id             string `json:"id"`
	OrganizationId string `json:"organization_id"`
	Server         string `json:"server"`
}

func (c *PritunlClient) LinkServerOrganizationGet(server string, organizationId string) (*LinkServerOrganizationData, error) {

	req := Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s/organization/%s", server, organizationId),
	}

	data := &LinkServerOrganizationData{}
	err := c.Do(req, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *PritunlClient) LinkServerOrganizationUpdate(server string, organizationId string, o LinkServerOrganizationPostData) (*LinkServerOrganizationData, error) {

	req := Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/organization/%s", server, organizationId),
		Json:   &o,
	}

	data := &LinkServerOrganizationData{}

	err := c.Do(req, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *PritunlClient) LinkServerOrganizationCreate(server string, organizationId string) (*LinkServerOrganizationData, error) {

	req := Request{
		Method: "POST",
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
