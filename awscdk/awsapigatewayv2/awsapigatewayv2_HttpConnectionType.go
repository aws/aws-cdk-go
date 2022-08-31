package awsapigatewayv2


// Supported connection types.
// Experimental.
type HttpConnectionType string

const (
	// For private connections between API Gateway and resources in a VPC.
	// Experimental.
	HttpConnectionType_VPC_LINK HttpConnectionType = "VPC_LINK"
	// For connections through public routable internet.
	// Experimental.
	HttpConnectionType_INTERNET HttpConnectionType = "INTERNET"
)

