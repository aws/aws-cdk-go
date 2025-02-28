package awsapigateway


// Example:
//   var integration lambdaIntegration
//   var resource resource
//   var responseModel model
//   var errorResponseModel model
//
//
//   resource.AddMethod(jsii.String("GET"), integration, &MethodOptions{
//   	// We can mark the parameters as required
//   	RequestParameters: map[string]*bool{
//   		"method.request.querystring.who": jsii.Boolean(true),
//   	},
//   	// we can set request validator options like below
//   	RequestValidatorOptions: &RequestValidatorOptions{
//   		RequestValidatorName: jsii.String("test-validator"),
//   		ValidateRequestBody: jsii.Boolean(true),
//   		ValidateRequestParameters: jsii.Boolean(false),
//   	},
//   	MethodResponses: []methodResponse{
//   		&methodResponse{
//   			// Successful response from the integration
//   			StatusCode: jsii.String("200"),
//   			// Define what parameters are allowed or not
//   			ResponseParameters: map[string]*bool{
//   				"method.response.header.Content-Type": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Origin": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Credentials": jsii.Boolean(true),
//   			},
//   			// Validate the schema on the response
//   			ResponseModels: map[string]iModel{
//   				"application/json": responseModel,
//   			},
//   		},
//   		&methodResponse{
//   			// Same thing for the error responses
//   			StatusCode: jsii.String("400"),
//   			ResponseParameters: map[string]*bool{
//   				"method.response.header.Content-Type": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Origin": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Credentials": jsii.Boolean(true),
//   			},
//   			ResponseModels: map[string]*iModel{
//   				"application/json": errorResponseModel,
//   			},
//   		},
//   	},
//   })
//
type RequestValidatorOptions struct {
	// The name of this request validator.
	// Default: None.
	//
	RequestValidatorName *string `field:"optional" json:"requestValidatorName" yaml:"requestValidatorName"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	// Default: false.
	//
	ValidateRequestBody *bool `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	// Default: false.
	//
	ValidateRequestParameters *bool `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

