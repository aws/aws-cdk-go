// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha


// Specifies the type responses the lambda returns.
// Experimental.
type HttpLambdaResponseType string

const (
	// Returns simple boolean response.
	// Experimental.
	HttpLambdaResponseType_SIMPLE HttpLambdaResponseType = "SIMPLE"
	// Returns an IAM Policy.
	// Experimental.
	HttpLambdaResponseType_IAM HttpLambdaResponseType = "IAM"
)

