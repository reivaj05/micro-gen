package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/reivaj05/GoJSON"
	"github.com/reivaj05/GoServer"
)

func getListHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement list")
	jsonResponse := createJSONListResponse()
	GoServer.SendResponseWithStatus(rw, jsonResponse, http.StatusOK)
}

func createJSONListResponse() string {
	json, _ := GoJSON.New("{}")
	json.CreateJSONArrayAtPath("data")
	return json.ToString()
}

func getItemHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement get")
	json, _ := GoJSON.New("{}")
	GoServer.SendResponseWithStatus(rw, json.ToString(), http.StatusOK)
}

func postItemhandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("TODO: Implement post")
	data, err := getJSONData(req)
	if err != nil {
		sendBadRequestResponse(rw)
		return
	}
	if !isDataValid(data) {
		sendBadRequestResponse(rw)
		return
	}
	json, _ := GoJSON.New("{}")
	GoServer.SendResponseWithStatus(rw, json.ToString(), http.StatusCreated)
}

func putItemHandler(rw http.ResponseWriter, req *http.Request) {
	data, err := getJSONData(req)
	if err != nil {
		sendBadRequestResponse(rw)
		return
	}
	params := GoServer.GetQueryParams(req)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("TODO: Implement put", id)
	if !isDataValid(data) {
		sendBadRequestResponse(rw)
		return
	}
	json, _ := GoJSON.New("{}")
	GoServer.SendResponseWithStatus(rw, json.ToString(), http.StatusOK)
}

func getJSONData(req *http.Request) (data *GoJSON.JSONWrapper, err error) {
	body, _ := GoServer.ReadBodyRequest(req)
	data, err = GoJSON.New(body)
	return
}

func isDataValid(data *GoJSON.JSONWrapper) bool {
	// TODO: Implement your own logic
	return true
}

func sendBadRequestResponse(rw http.ResponseWriter) {
	GoServer.SendResponseWithStatus(
		rw, GoServer.BadRequest, http.StatusBadRequest)
}

func deleteItemHandler(rw http.ResponseWriter, req *http.Request) {
	params := GoServer.GetQueryParams(req)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("TODO: Implement delete", id)
	GoServer.SendResponseWithStatus(rw, "", http.StatusOK)
}
