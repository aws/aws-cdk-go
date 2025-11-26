package awsapigateway


// The response transfer mode of the integration.
//
// Example:
//   var handler Function
//
//   apigateway.NewLambdaIntegration(handler, &LambdaIntegrationOptions{
//   	ResponseTransferMode: apigateway.ResponseTransferMode_STREAM,
//   })
//
type ResponseTransferMode string

const (
	// API Gateway waits to receive the complete response before beginning transmission.
	ResponseTransferMode_BUFFERED ResponseTransferMode = "BUFFERED"
	// API Gateway streams the response back to you as it is received from the integration.
	//
	// This is only supported for AWS_PROXY and HTTP_PROXY integration types.
	ResponseTransferMode_STREAM ResponseTransferMode = "STREAM"
)

