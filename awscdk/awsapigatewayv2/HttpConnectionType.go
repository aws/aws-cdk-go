package awsapigatewayv2


// Supported connection types.
type HttpConnectionType string

const (
	// For private connections between API Gateway and resources in a VPC.
	HttpConnectionType_VPC_LINK HttpConnectionType = "VPC_LINK"
	// For connections through public routable internet.
	HttpConnectionType_INTERNET HttpConnectionType = "INTERNET"
)

