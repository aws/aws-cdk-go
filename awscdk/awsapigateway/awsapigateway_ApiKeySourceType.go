package awsapigateway


type ApiKeySourceType string

const (
	// To read the API key from the `X-API-Key` header of a request.
	ApiKeySourceType_HEADER ApiKeySourceType = "HEADER"
	// To read the API key from the `UsageIdentifierKey` from a custom authorizer.
	ApiKeySourceType_AUTHORIZER ApiKeySourceType = "AUTHORIZER"
)

