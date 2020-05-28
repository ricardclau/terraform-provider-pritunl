package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type Server struct {
	Id      string
	Name    string
	Network string
	Groups  string
	Port    string
}

func LoadServer(d *schema.ResourceData) (sch *Server) {
	sch = &Server{
		Id:   d.Id(),
		Name: d.Get("name").(string),
	}

	return
}
