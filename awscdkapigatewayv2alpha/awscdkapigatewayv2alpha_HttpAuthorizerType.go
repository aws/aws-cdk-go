// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Supported Authorizer types.
// Experimental.
type HttpAuthorizerType string

const (
	// IAM Authorizer.
	// Experimental.
	HttpAuthorizerType_IAM HttpAuthorizerType = "IAM"
	// JSON Web Tokens.
	// Experimental.
	HttpAuthorizerType_JWT HttpAuthorizerType = "JWT"
	// Lambda Authorizer.
	// Experimental.
	HttpAuthorizerType_LAMBDA HttpAuthorizerType = "LAMBDA"
)

