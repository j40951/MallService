package service

import (
	"github.com/emicklei/go-restful"
	"github.com/j40951/MallService/delegate"
)

// MallService means Mall Web Service
func MallService(delegater delegate.MallDelegater) *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/mall/v1")
	ws.Route(ws.GET("/hello").To(delegater.Hello))
	return ws
}
