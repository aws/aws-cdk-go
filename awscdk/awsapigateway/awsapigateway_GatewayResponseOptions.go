package awsapigateway


// Options to add gateway response.
//
// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books-api"))
//   api.addGatewayResponse(jsii.String("test-response"), &gatewayResponseOptions{
//   	type: apigateway.responseType_ACCESS_DENIED(),
//   	statusCode: jsii.String("500"),
//   	responseHeaders: map[string]*string{
//   		// Note that values must be enclosed within a pair of single quotes
//   		"Access-Control-Allow-Origin": jsii.String("'test.com'"),
//   		"test-key": jsii.String("'test-value'"),
//   	},
//   	templates: map[string]*string{
//   		"application/json": jsii.String("{ \"message\": $context.error.messageString, \"statusCode\": \"488\", \"type\": \"$context.error.responseType\" }"),
//   	},
//   })
//
type GatewayResponseOptions struct {
	// Response type to associate with gateway response.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html
	//
	Type ResponseType `field:"required" json:"type" yaml:"type"`
	// Custom headers parameters for response.
	ResponseHeaders *map[string]*string `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
	// Http status code for response.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Custom templates to get mapped as response.
	Templates *map[string]*string `field:"optional" json:"templates" yaml:"templates"`
}

