package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
)

/*  type Endpoints struct {
	GetRoles   endpoint.Endpoint
	ParseToken endpoint.Endpoint
}  */

type Endpoints struct {
	ParseTokenEndpoint endpoint.Endpoint
	GetRolesEndpoint   endpoint.Endpoint
}

type ParseTokenRequest struct {
	Token string `json:"token"`
}

type ParseTokenResponse struct {
	UserId int `json:"user_id"`
	//Err    string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type GetRolesRequest struct {
	UserId int `json:"user_id"`
}

type GetRolesResponse struct {
	Roles []string `json:"roles"`
}

func MakeParseTokenEndpoint(a *authService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(ParseTokenRequest)

		userId, err := a.ParseToken(ctx, req.Token)
		if err != nil {
			return ParseTokenResponse{userId}, err
		}

		return ParseTokenResponse{userId}, nil
	}
}

func MakeGetRolesEndpoint(a *authService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetRolesRequest)
		roles, err := a.GetRoles(ctx, req.UserId)
		if err != nil {
			return GetRolesResponse{roles}, err
		}

		return GetRolesResponse{roles}, nil
	}
}

func decodeGetRolesResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp GetRolesResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeParseTokenResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp ParseTokenResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func encodeGetRolesRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.ContentLength = int64(len(buf.Bytes()))
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeParseTokenRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.ContentLength = int64(len(buf.Bytes()))
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func copyUrl(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
