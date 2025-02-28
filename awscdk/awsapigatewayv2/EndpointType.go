package awsapigatewayv2


// Endpoint type for a domain name.
type EndpointType string

const (
	// For an edge-optimized custom domain name.
	EndpointType_EDGE EndpointType = "EDGE"
	// For a regional custom domain name.
	EndpointType_REGIONAL EndpointType = "REGIONAL"
)

