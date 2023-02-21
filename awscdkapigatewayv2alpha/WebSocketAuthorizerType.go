// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Supported Authorizer types.
// Experimental.
type WebSocketAuthorizerType string

const (
	// Lambda Authorizer.
	// Experimental.
	WebSocketAuthorizerType_LAMBDA WebSocketAuthorizerType = "LAMBDA"
	// IAM Authorizer.
	// Experimental.
	WebSocketAuthorizerType_IAM WebSocketAuthorizerType = "IAM"
)

