package awsbedrockagentcorealpha


// Configuration for API Gateway tools.
//
// The API Gateway tool configuration defines which operations from your REST API
// are exposed as tools. It requires a list of tool filters to select operations
// to expose, and optionally accepts tool overrides to customize tool metadata.
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
//   // Uses IAM authorization for outbound auth by default
//   apiGatewayTarget := gateway.AddApiGatewayTarget(jsii.String("MyApiGatewayTarget"), &AddApiGatewayTargetOptions{
//   	RestApi: api,
//   	ApiGatewayToolConfiguration: &ApiGatewayToolConfiguration{
//   		ToolFilters: []ApiGatewayToolFilter{
//   			&ApiGatewayToolFilter{
//   				FilterPath: jsii.String("/pets/*"),
//   				Methods: []ApiGatewayHttpMethod{
//   					agentcore.ApiGatewayHttpMethod_GET,
//   				},
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type ApiGatewayToolConfiguration struct {
	// List of tool filters to select operations At least one filter is required.
	// Experimental.
	ToolFilters *[]*ApiGatewayToolFilter `field:"required" json:"toolFilters" yaml:"toolFilters"`
	// Optional list of tool overrides to customize tool metadata Each override must correspond to an operation selected by the filters.
	// Default: - No tool overrides.
	//
	// Experimental.
	ToolOverrides *[]*ApiGatewayToolOverride `field:"optional" json:"toolOverrides" yaml:"toolOverrides"`
}

