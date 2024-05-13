package constants

// application constants
const (
	Empty                = ""
	ContextAppNameKey    = "appName"
	TokenExpirationHours = 12
	ApiKeyLength         = 24

	RootUserEmailEnvKey    = "ROOT_USER_EMAIL"
	RootUserPasswordEnvKey = "ROOT_USER_PASSWORD"

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
	CreateAppRoute    = "/createapp"

	// app route constants
	RedirectUrlRoute   = "/:slug"
	RedirectToAppRoute = "/"
)
