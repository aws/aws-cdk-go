package awsapigateway


// Experimental.
type ApiKeySourceType string

const (
	// To read the API key from the `X-API-Key` header of a request.
	// Experimental.
	ApiKeySourceType_HEADER ApiKeySourceType = "HEADER"
	// To read the API key from the `UsageIdentifierKey` from a custom authorizer.
	// Experimental.
	ApiKeySourceType_AUTHORIZER ApiKeySourceType = "AUTHORIZER"
)

