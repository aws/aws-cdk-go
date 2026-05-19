package awsbedrockagentcorealpha


// Configuration for API Gateway tools.
//
// The API Gateway tool configuration defines which operations from your REST API
// are exposed as tools. It requires a list of tool filters to select operations
// to expose, and optionally accepts tool overrides to customize tool metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   apiGatewayToolConfiguration := &ApiGatewayToolConfiguration{
//   	ToolFilters: []ApiGatewayToolFilter{
//   		&ApiGatewayToolFilter{
//   			FilterPath: jsii.String("filterPath"),
//   			Methods: []ApiGatewayHttpMethod{
//   				bedrock_agentcore_alpha.ApiGatewayHttpMethod_GET,
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ToolOverrides: []ApiGatewayToolOverride{
//   		&ApiGatewayToolOverride{
//   			Method: bedrock_agentcore_alpha.ApiGatewayHttpMethod_GET,
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   }
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

