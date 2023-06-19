package awsapigateway


// Options to add gateway response.
//
// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books-api"))
//   api.AddGatewayResponse(jsii.String("test-response"), &GatewayResponseOptions{
//   	Type: apigateway.ResponseType_ACCESS_DENIED(),
//   	StatusCode: jsii.String("500"),
//   	ResponseHeaders: map[string]*string{
//   		"Access-Control-Allow-Origin": jsii.String("test.com"),
//   		"test-key": jsii.String("test-value"),
//   	},
//   	Templates: map[string]*string{
//   		"application/json": jsii.String("{ \"message\": $context.error.messageString, \"statusCode\": \"488\", \"type\": \"$context.error.responseType\" }"),
//   	},
//   })
//
// Experimental.
type GatewayResponseOptions struct {
	// Response type to associate with gateway response.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html
	//
	// Experimental.
	Type ResponseType `field:"required" json:"type" yaml:"type"`
	// Custom headers parameters for response.
	// Experimental.
	ResponseHeaders *map[string]*string `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
	// Http status code for response.
	// Experimental.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Custom templates to get mapped as response.
	// Experimental.
	Templates *map[string]*string `field:"optional" json:"templates" yaml:"templates"`
}

