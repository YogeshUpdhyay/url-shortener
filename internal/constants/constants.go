package constants

// application constants
const (
	Empty             = ""
	ContextAppNameKey = "appName"

	// header keys
	AuthHeaderKey   = "Authorization"
	ApiKeyHeaderKey = "X-Api-Key"
)

// routes
const (
	ApiRoute = "/api"
	V1Route  = "/v1"

	// url route constants
	ShortenUrlRoute   = "/shortenurl"
	DeleteUrlRoute    = "/deleteurl"
	AuthenticateRoute = "/authenticate"

	// app route constants
	RedirectUrlRoute   = "/:slug"
	RedirectToAppRoute = "/"
)
