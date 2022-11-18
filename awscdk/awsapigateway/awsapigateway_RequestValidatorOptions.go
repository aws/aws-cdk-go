package awsapigateway


// Example:
//   var integration lambdaIntegration
//   var resource resource
//   var responseModel model
//   var errorResponseModel model
//
//
//   resource.addMethod(jsii.String("GET"), integration, &methodOptions{
//   	// We can mark the parameters as required
//   	requestParameters: map[string]*bool{
//   		"method.request.querystring.who": jsii.Boolean(true),
//   	},
//   	// we can set request validator options like below
//   	requestValidatorOptions: &requestValidatorOptions{
//   		requestValidatorName: jsii.String("test-validator"),
//   		validateRequestBody: jsii.Boolean(true),
//   		validateRequestParameters: jsii.Boolean(false),
//   	},
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			// Successful response from the integration
//   			statusCode: jsii.String("200"),
//   			// Define what parameters are allowed or not
//   			responseParameters: map[string]*bool{
//   				"method.response.header.Content-Type": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Origin": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Credentials": jsii.Boolean(true),
//   			},
//   			// Validate the schema on the response
//   			responseModels: map[string]iModel{
//   				"application/json": responseModel,
//   			},
//   		},
//   		&methodResponse{
//   			// Same thing for the error responses
//   			statusCode: jsii.String("400"),
//   			responseParameters: map[string]*bool{
//   				"method.response.header.Content-Type": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Origin": jsii.Boolean(true),
//   				"method.response.header.Access-Control-Allow-Credentials": jsii.Boolean(true),
//   			},
//   			responseModels: map[string]*iModel{
//   				"application/json": errorResponseModel,
//   			},
//   		},
//   	},
//   })
//
type RequestValidatorOptions struct {
	// The name of this request validator.
	RequestValidatorName *string `field:"optional" json:"requestValidatorName" yaml:"requestValidatorName"`
	// Indicates whether to validate the request body according to the configured schema for the targeted API and method.
	ValidateRequestBody *bool `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// Indicates whether to validate request parameters.
	ValidateRequestParameters *bool `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

