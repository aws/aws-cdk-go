package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for creating a Lambda-based Gateway Target Convenience interface for the most common use case.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_22_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	        exports.handler = async (event) => {
//   	            return {
//   	                statusCode: 200,
//   	                body: JSON.stringify({ message: 'Hello from Lambda!' })
//   	            };
//   	        };
//   	    `)),
//   })
//
//   // Create a gateway target with Lambda and tool schema
//   target := agentcore.GatewayTarget_ForLambda(this, jsii.String("MyLambdaTarget"), &GatewayTargetLambdaProps{
//   	GatewayTargetName: jsii.String("my-lambda-target"),
//   	Description: jsii.String("Target for Lambda function integration"),
//   	Gateway: gateway,
//   	LambdaFunction: lambdaFunction,
//   	ToolSchema: agentcore.ToolSchema_FromLocalAsset(path.join(__dirname, jsii.String("schemas"), jsii.String("my-tool-schema.json"))),
//   })
//
// Experimental.
type GatewayTargetLambdaProps struct {
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Experimental.
	GatewayTargetName *string `field:"required" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
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

