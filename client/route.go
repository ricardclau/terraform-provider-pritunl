package client

import (
	"fmt"
	"log"
)

type RouteData struct {
	Id           string `json:"id"`
	Network      string `json:"network"`
	Comment      string `json:"comment"`
	Metric       string `json:"metric"`
	Nat          bool   `json:"nat"`
	NatInterface string `json:"nat_interface"`
	NatNetmap    string `json:"nat_netmap"`
	NetGateway   bool   `json:"net_gateway"`
	Advertise    bool   `json:"advertise"`
}

type RoutePostData struct {
	Network      string `json:"network"`
	Comment      string `json:"comment"`
	Metric       string `json:"metric,omitempty"`
	Nat          bool   `json:"nat"`
	NatInterface string `json:"nat_interface"`
	NatNetmap    string `json:"nat_netmap"`
	NetGateway   bool   `json:"net_gateway"`
	Advertise    bool   `json:"advertise"`
}

func (c *PritunlClient) RouteGetByNetwork(serverId string, network string) (*RouteData, error) {
	req := Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s/route", serverId),
	}

	var data []RouteData
	err := c.Do(req, &data)
	if err != nil {
		return nil, err
	}

	for _, route := range data {
		log.Println(fmt.Sprintf("[DEBUG] route.Network: %s vs network %s", route.Network, network))
		if route.Network == network {
			return &route, nil
		}
	}

	return nil, fmt.Errorf("Cannot find route with network: %v in server: %v", network, serverId)
}

func (c *PritunlClient) RouteGet(serverId string, routeId string) (*RouteData, error) {
	req := Request{
		Method: "GET",
		Path:   fmt.Sprintf("/server/%s/route", serverId),
	}

	var data []RouteData
	err := c.Do(req, &data)
	if err != nil {
		return nil, err
	}

	for _, route := range data {
		log.Println(fmt.Sprintf("[DEBUG] route.Id: %s vs routeId %s", route.Id, routeId))
		if route.Id == routeId {
			return &route, nil
		}
	}

	return nil, fmt.Errorf("Cannot find route with id: %v in server: %v", routeId, serverId)
}

func (c *PritunlClient) RouteCreate(serverId string, r RoutePostData) (*RouteData, error) {
	req := Request{
		Method: "POST",
		Path:   fmt.Sprintf("/server/%s/route", serverId),
		Json:   r,
	}

	data := &RouteData{}
	err := c.Do(req, data)

	return data, err
}

func (c *PritunlClient) RouteUpdate(serverId string, routeId string, r RoutePostData) (*RouteData, error) {
	req := Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/server/%s/route/%s", serverId, routeId),
		Json:   r,
	}

	data := &RouteData{}
	err := c.Do(req, data)

	return data, err
}

func (c *PritunlClient) RouteDelete(serverId string, routeId string) error {
	req := Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/server/%s/route/%s", serverId, routeId),
	}

	return c.Do(req, nil)
}
