package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Properties for creating an API Gateway target configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var restApi RestApi
//
//   apiGatewayTargetConfigurationProps := &ApiGatewayTargetConfigurationProps{
//   	ApiGatewayToolConfiguration: &ApiGatewayToolConfiguration{
//   		ToolFilters: []ApiGatewayToolFilter{
//   			&ApiGatewayToolFilter{
//   				FilterPath: jsii.String("filterPath"),
//   				Methods: []ApiGatewayHttpMethod{
//   					bedrock_agentcore_alpha.ApiGatewayHttpMethod_GET,
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		ToolOverrides: []ApiGatewayToolOverride{
//   			&ApiGatewayToolOverride{
//   				Method: bedrock_agentcore_alpha.ApiGatewayHttpMethod_GET,
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	RestApi: restApi,
//
//   	// the properties below are optional
//   	MetadataConfiguration: &MetadataConfiguration{
//   		AllowedQueryParameters: []*string{
//   			jsii.String("allowedQueryParameters"),
//   		},
//   		AllowedRequestHeaders: []*string{
//   			jsii.String("allowedRequestHeaders"),
//   		},
//   		AllowedResponseHeaders: []*string{
//   			jsii.String("allowedResponseHeaders"),
//   		},
//   	},
//   	Stage: jsii.String("stage"),
//   }
//
// Experimental.
type ApiGatewayTargetConfigurationProps struct {
	// Tool configuration defining which operations to expose.
	// Experimental.
	ApiGatewayToolConfiguration *ApiGatewayToolConfiguration `field:"required" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
	// The REST API to integrate with Must be in the same account and region as the gateway.
	// Experimental.
	RestApi awsapigateway.IRestApi `field:"required" json:"restApi" yaml:"restApi"`
	// Metadata configuration for passing headers and query parameters Allows you to pass additional context through headers and query parameters.
	// Default: - No metadata configuration.
	//
	// Experimental.
	MetadataConfiguration *MetadataConfiguration `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// The stage name of the REST API.
	// Default: - Uses the deployment stage from the RestApi (restApi.deploymentStage.stageName)
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

