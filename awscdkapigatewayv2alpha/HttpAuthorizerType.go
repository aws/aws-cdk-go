package awscdkapigatewayv2alpha


// Supported Authorizer types.
// Deprecated.
type HttpAuthorizerType string

const (
	// IAM Authorizer.
	// Deprecated.
	HttpAuthorizerType_IAM HttpAuthorizerType = "IAM"
	// JSON Web Tokens.
	// Deprecated.
	HttpAuthorizerType_JWT HttpAuthorizerType = "JWT"
	// Lambda Authorizer.
	// Deprecated.
	HttpAuthorizerType_LAMBDA HttpAuthorizerType = "LAMBDA"
)

