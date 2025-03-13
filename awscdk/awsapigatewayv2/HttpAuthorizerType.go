package awsapigatewayv2


// Supported Authorizer types.
type HttpAuthorizerType string

const (
	// IAM Authorizer.
	HttpAuthorizerType_IAM HttpAuthorizerType = "IAM"
	// JSON Web Tokens.
	HttpAuthorizerType_JWT HttpAuthorizerType = "JWT"
	// Lambda Authorizer.
	HttpAuthorizerType_LAMBDA HttpAuthorizerType = "LAMBDA"
)

