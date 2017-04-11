package main

// ********* Modify this file accordingly to your needs ************

import (
	"github.com/reivaj05/GoServer"
)

var Endpoints = []*GoServer.Endpoint{
	&GoServer.Endpoint{
		Method: "GET",
		Path:    "/{id:[0-9]+}",
		Handler: getItemHandler,
	},
	&GoServer.Endpoint{
		Method:  "GET",
		Path:    "/",
		Handler: getListHandler,
	},
	&GoServer.Endpoint{
		Method:  "POST",
		Path:    "/",
		Handler: postItemhandler,
	},
	&GoServer.Endpoint{
		Method:  "PUT",
		Path:    "/{id:[0-9]+}",
		Handler: putItemHandler,
	},
	&GoServer.Endpoint{
		Method:  "DELETE",
		Path:    "/{id:[0-9]+}",
		Handler: deleteItemHandler,
	},
}
