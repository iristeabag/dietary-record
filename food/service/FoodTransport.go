package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DecodeGetFoodByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetFoodByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetFoodByIdRequest{
		Id: vars["foodid"],
	}

	return req, nil
}

func DecodeGetAllFoodsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into GETALL Decoding")
	var req GetAllFoodsRequest
	return req, nil
}

func DecodeCreateFoodRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateFoodRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.food); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdateFoodRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateFoodRequest
	if err := json.NewDecoder(r.Body).Decode(&req.food); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

func DecodeDeleteFoodRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteFoodRequest
	vars := mux.Vars(r)
	req = DeleteFoodRequest{
		Foodid: vars["foodid"],
	}
	return req, nil
}
