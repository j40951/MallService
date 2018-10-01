package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"

	"github.com/j40951/MallService/resource"
	"github.com/j40951/MallService/service"
)

func main() {
	fmt.Print("hello world")
	mallResource := new(resource.MallResource)
	restful.Add(service.MallService(*mallResource))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
