package resource

import (
	"io"

	"github.com/emicklei/go-restful"
)

// MallResource means Mall Resource
type MallResource struct {
}

// Hello measn /mall/hello
func (m MallResource) Hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "Mall")
}
