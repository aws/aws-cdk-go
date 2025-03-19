package awsapigatewayv2authorizers


// Specifies the type responses the lambda returns.
type HttpLambdaResponseType string

const (
	// Returns simple boolean response.
	HttpLambdaResponseType_SIMPLE HttpLambdaResponseType = "SIMPLE"
	// Returns an IAM Policy.
	HttpLambdaResponseType_IAM HttpLambdaResponseType = "IAM"
)

