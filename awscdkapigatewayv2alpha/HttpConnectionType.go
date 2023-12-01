package awscdkapigatewayv2alpha


// Supported connection types.
// Deprecated.
type HttpConnectionType string

const (
	// For private connections between API Gateway and resources in a VPC.
	// Deprecated.
	HttpConnectionType_VPC_LINK HttpConnectionType = "VPC_LINK"
	// For connections through public routable internet.
	// Deprecated.
	HttpConnectionType_INTERNET HttpConnectionType = "INTERNET"
)

