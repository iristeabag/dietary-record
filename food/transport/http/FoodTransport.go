package service

import (
	"context"
	"encoding/json"
	"fmt"
	endpoint "go-kit-demo/food/endpoint/http"
	svc "go-kit-demo/food/service"
	"net/http"

	"github.com/gorilla/mux"
)

func DecodeGetFoodByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.GetFoodByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = endpoint.GetFoodByIdRequest{
		Id: vars["foodid"],
	}

	return req, nil
}

func DecodeGetAllFoodsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into GETALL Decoding")
	var req endpoint.GetAllFoodsRequest
	return req, nil
}

func DecodeCreateFoodRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var req endpoint.CreateFoodRequest
	var req svc.Food
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdateFoodRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	// var req endpoint.UpdateFoodRequest
	vars := mux.Vars(r)
	var req svc.Food
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.Foodid = vars["foodid"]
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

func DecodeDeleteFoodRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req endpoint.DeleteFoodRequest
	vars := mux.Vars(r)
	req = endpoint.DeleteFoodRequest{
		Foodid: vars["foodid"],
	}
	return req, nil
}
