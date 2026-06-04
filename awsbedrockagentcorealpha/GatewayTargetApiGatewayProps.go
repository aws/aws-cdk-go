package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Properties for creating an API Gateway-based Gateway Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var credentialProviderConfig ICredentialProviderConfig
//   var gateway Gateway
//   var restApi RestApi
//
//   gatewayTargetApiGatewayProps := &GatewayTargetApiGatewayProps{
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
//   	Gateway: gateway,
//   	RestApi: restApi,
//
//   	// the properties below are optional
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		credentialProviderConfig,
//   	},
//   	Description: jsii.String("description"),
//   	GatewayTargetName: jsii.String("gatewayTargetName"),
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
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type GatewayTargetApiGatewayProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Tool configuration defining which operations to expose.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ApiGatewayToolConfiguration *ApiGatewayToolConfiguration `field:"required" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
	// The gateway this target belongs to [disable-awslint:prefer-ref-interface].
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// The REST API to integrate with Must be in the same account and region as the gateway [disable-awslint:prefer-ref-interface].
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	RestApi awsapigateway.IRestApi `field:"required" json:"restApi" yaml:"restApi"`
	// Credential providers for authentication API Gateway targets support IAM and API key authentication.
	// Default: - Empty array (service handles IAM automatically).
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Metadata configuration for passing headers and query parameters Allows you to pass additional context through headers and query parameters.
	// Default: - No metadata configuration.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MetadataConfiguration *MetadataConfiguration `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// The stage name of the REST API.
	// Default: - Uses the deployment stage from the RestApi (restApi.deploymentStage.stageName)
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

