package awsapigateway


// Example:
//   var getBookHandler function
//   var getBookIntegration lambdaIntegration
//
//
//   getBookIntegration := apigateway.NewLambdaIntegration(getBookHandler, &lambdaIntegrationOptions{
//   	contentHandling: apigateway.contentHandling_CONVERT_TO_TEXT,
//   	 // convert to base64
//   	credentialsPassthrough: jsii.Boolean(true),
//   })
//
type ContentHandling string

const (
	// Converts a request payload from a base64-encoded string to a binary blob.
	ContentHandling_CONVERT_TO_BINARY ContentHandling = "CONVERT_TO_BINARY"
	// Converts a request payload from a binary blob to a base64-encoded string.
	ContentHandling_CONVERT_TO_TEXT ContentHandling = "CONVERT_TO_TEXT"
)

