package awsapigatewayv2


// Supported Authorizer types.
type WebSocketAuthorizerType string

const (
	// Lambda Authorizer.
	WebSocketAuthorizerType_LAMBDA WebSocketAuthorizerType = "LAMBDA"
	// IAM Authorizer.
	WebSocketAuthorizerType_IAM WebSocketAuthorizerType = "IAM"
)

