package awsbedrockagentcorealpha


// Configuration for passing metadata (headers and query parameters) to the API Gateway target.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("MyApi"), &RestApiProps{
//   	RestApiName: jsii.String("my-api"),
//   })
//
//   // Create a gateway target using the static factory method
//   apiGatewayTarget := agentcore.GatewayTarget_ForApiGateway(this, jsii.String("MyApiGatewayTarget"), &GatewayTargetApiGatewayProps{
//   	GatewayTargetName: jsii.String("my-api-gateway-target"),
//   	Description: jsii.String("Target for API Gateway REST API integration"),
//   	Gateway: gateway,
//   	RestApi: api,
//   	ApiGatewayToolConfiguration: &ApiGatewayToolConfiguration{
//   		ToolFilters: []ApiGatewayToolFilter{
//   			&ApiGatewayToolFilter{
//   				FilterPath: jsii.String("/pets/*"),
//   				Methods: []ApiGatewayHttpMethod{
//   					agentcore.ApiGatewayHttpMethod_GET,
//   					agentcore.ApiGatewayHttpMethod_POST,
//   				},
//   			},
//   		},
//   	},
//   	MetadataConfiguration: &MetadataConfiguration{
//   		AllowedRequestHeaders: []*string{
//   			jsii.String("X-User-Id"),
//   		},
//   		AllowedQueryParameters: []*string{
//   			jsii.String("limit"),
//   		},
//   	},
//   })
//
// Experimental.
type MetadataConfiguration struct {
	// List of query parameter names to pass through to the target.
	//
	// Constraints:
	// - Array must contain 1-10 items
	// - Each parameter name must be 1-40 characters
	// - Cannot be an empty array.
	// Default: - No query parameters are passed through.
	//
	// Experimental.
	AllowedQueryParameters *[]*string `field:"optional" json:"allowedQueryParameters" yaml:"allowedQueryParameters"`
	// List of request header names to pass through to the target.
	//
	// Constraints:
	// - Array must contain 1-10 items
	// - Each header name must be 1-100 characters
	// - Cannot be an empty array.
	// Default: - No request headers are passed through.
	//
	// Experimental.
	AllowedRequestHeaders *[]*string `field:"optional" json:"allowedRequestHeaders" yaml:"allowedRequestHeaders"`
	// List of response header names to pass through from the target.
	//
	// Constraints:
	// - Array must contain 1-10 items
	// - Each header name must be 1-100 characters
	// - Cannot be an empty array.
	// Default: - No response headers are passed through.
	//
	// Experimental.
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
}

