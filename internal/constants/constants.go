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
	AuthHeaderKey               = "Authorization"
	ApiKeyHeaderKey             = "X-Api-Key"
	AuthenticatedUserContextKey = "USER"

	// pagination
	PaginationPageNumberKey = "pageNumber"
	PaginationPageSizeKey   = "pageSize"
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
	ListAppsRoute     = "/listapps"

	// app route constants
	RedirectUrlRoute   = "/:slug"
	RedirectToAppRoute = "/"
)
