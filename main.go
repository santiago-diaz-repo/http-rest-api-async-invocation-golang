package main

import (
	"http-rest-api-versioning-golang/handler"
)

func main() {
	exampleHandler := &handler.Handler{}
	exampleHandler.HandleRequest()
}
