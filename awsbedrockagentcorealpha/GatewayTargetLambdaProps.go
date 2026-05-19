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
// Experimental.
type GatewayTargetLambdaProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// The gateway this target belongs to.
	// Experimental.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
	// The Lambda function to associate with this target.
	// Experimental.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// The tool schema defining the available tools.
	// Experimental.
	ToolSchema ToolSchema `field:"required" json:"toolSchema" yaml:"toolSchema"`
	// Credential providers for authentication Lambda targets only support IAM role authentication.
	// Default: - [GatewayCredentialProvider.fromIamRole()]
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
}

