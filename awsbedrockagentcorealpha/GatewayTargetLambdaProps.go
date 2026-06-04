package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for creating a Lambda-based Gateway Target Convenience interface for the most common use case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var credentialProviderConfig ICredentialProviderConfig
//   var function_ Function
//   var gateway Gateway
//   var toolSchema ToolSchema
//
//   gatewayTargetLambdaProps := &GatewayTargetLambdaProps{
//   	Gateway: gateway,
//   	LambdaFunction: function_,
//   	ToolSchema: toolSchema,
//
//   	// the properties below are optional
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		credentialProviderConfig,
//   	},
//   	Description: jsii.String("description"),
//   	GatewayTargetName: jsii.String("gatewayTargetName"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type GatewayTargetLambdaProps struct {
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
	// The gateway this target belongs to.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// The Lambda function to associate with this target.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// The tool schema defining the available tools.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ToolSchema ToolSchema `field:"required" json:"toolSchema" yaml:"toolSchema"`
	// Credential providers for authentication Lambda targets only support IAM role authentication.
	// Default: - [GatewayCredentialProvider.fromIamRole()]
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
}

