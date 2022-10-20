// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Endpoint type for a domain name.
// Experimental.
type EndpointType string

const (
	// For an edge-optimized custom domain name.
	// Experimental.
	EndpointType_EDGE EndpointType = "EDGE"
	// For a regional custom domain name.
	// Experimental.
	EndpointType_REGIONAL EndpointType = "REGIONAL"
)

