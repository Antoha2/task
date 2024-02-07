package web

import (
	"context"
	"encoding/json"

	"log"
	"net/http"

	taskEndpoints "github.com/antoha2/task/transport/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func (wImpl *webImpl) StartHTTP() error {

	TaskOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(wImpl.UserIdentify),
	}

	CreateHandler := httptransport.NewServer(
		taskEndpoints.MakeCreateEndpoint(wImpl.taskService), //use the endpoint
		decodeMakeCreateRequest,                             //converts the parameters received via the request body into the struct expected by the endpoint
		encodeResponse,
		TaskOptions...,
	)

	ReadHandler := httptransport.NewServer(
		taskEndpoints.MakeReadEndpoint(wImpl.taskService),
		decodeMakeReadRequest,
		encodeResponse,
		TaskOptions...,
	)

	UpdateHandler := httptransport.NewServer(
		taskEndpoints.MakeUpdateEndpoint(wImpl.taskService),
		decodeMakeUpdateRequest,
		encodeResponse,
		TaskOptions...,
	)

	DeleteHandler := httptransport.NewServer(
		taskEndpoints.MakeDeleteEndpoint(wImpl.taskService),
		decodeMakeDeleteRequest,
		encodeResponse,
		TaskOptions...,
	)

	r := mux.NewRouter() //I'm using Gorilla Mux, but it could be any other library, or even the stdlib
	r.Methods("POST").Path("/api/create").Handler(CreateHandler)
	r.Methods("POST").Path("/api/read").Handler(ReadHandler)
	r.Methods("POST").Path("/api/update").Handler(UpdateHandler)
	r.Methods("POST").Path("/api/delete").Handler(DeleteHandler)

	wImpl.server = &http.Server{Addr: ":8181"}
	log.Printf("(task) Запуск HTTP-сервера на http://127.0.0.1%s\n", wImpl.server.Addr) //:8181

	if err := http.ListenAndServe(wImpl.server.Addr, r); err != nil {
		log.Println(err)
	}

	return nil
}

func StartGRPC() (*grpc.ClientConn, error) {
	var err error
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	GRPCConn, err := grpc.Dial("auth:8183", opts...)
	if err != nil {
		grpclog.Fatalf("task.StartGRPC() - fail to dial: %v", err)
		return nil, err
	}
	log.Println("(task) установка соединения с GRPC-сервером на http://127.0.0.1:8183") //:8183

	return GRPCConn, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func decodeMakeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request taskEndpoints.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	log.Println(r.Body)
	return request, nil
}

func decodeMakeReadRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request taskEndpoints.ReadRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request taskEndpoints.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request taskEndpoints.DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
