package delegate

import (
	"github.com/emicklei/go-restful"
)

// MallDelegater delegate MallService
type MallDelegater interface {
	Hello(req *restful.Request, resp *restful.Response)
}
