package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/greeting",
		Index,
	},
	Route{
		"TestApi",
		"POST",
		"/testApi",
		TestApi,
	},
	Route{
		"GetAuthUrl",
		"POST",
		"/getAuthUrl",
		GetAuthUrl,
	},
	Route{
		"GetCtx",
		"GET",
		"/getctx",
		GetCtx,
	},

	Route{
		"OauthCallback",
		"GET",
		"/oauth2callback",
		OAuthCallback,
	},
}
