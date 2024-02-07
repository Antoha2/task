package auth

import (
	"log"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	//"github.com/go-kit/log"
)

func NewEndpoints(Url string) *Endpoints {
	addr, err := url.Parse(Url)
	if err != nil {
		log.Println(err)
	}
	var options []httptransport.ClientOption

	var getRolesEndpoint endpoint.Endpoint
	{
		getRolesEndpoint = httptransport.NewClient(
			"POST",
			copyUrl(addr, "/auth/getRoles"),
			encodeGetRolesRequest,
			decodeGetRolesResponse,
			options...,
		).Endpoint()
	}

	var parseTokenEndpoint endpoint.Endpoint
	{
		parseTokenEndpoint = httptransport.NewClient(
			"POST",
			copyUrl(addr, "/auth/parseToken"),
			encodeParseTokenRequest,
			decodeParseTokenResponse,
			options...,
		).Endpoint()
	}

	return &Endpoints{

		GetRolesEndpoint:   getRolesEndpoint,
		ParseTokenEndpoint: parseTokenEndpoint,
	}
}
