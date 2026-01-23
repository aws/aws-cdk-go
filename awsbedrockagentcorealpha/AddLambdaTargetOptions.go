package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Options for adding a Lambda target to a gateway.
//
// Example:
//   // Create a gateway first
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_22_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	    exports.handler = async (event) => {
//   	      return {
//   	        statusCode: 200,
//   	        body: JSON.stringify({ message: 'Hello from Lambda!' })
//   	      };
//   	    };
//   	  `)),
//   })
//
//   lambdaTarget := gateway.AddLambdaTarget(jsii.String("MyLambdaTarget"), &AddLambdaTargetOptions{
//   	GatewayTargetName: jsii.String("my-lambda-target"),
//   	Description: jsii.String("Lambda function target"),
//   	LambdaFunction: lambdaFunction,
//   	ToolSchema: agentcore.ToolSchema_FromInline([]ToolDefinition{
//   		&ToolDefinition{
//   			Name: jsii.String("hello_world"),
//   			Description: jsii.String("A simple hello world tool"),
//   			InputSchema: &SchemaDefinition{
//   				Type: agentcore.SchemaDefinitionType_OBJECT,
//   				Properties: map[string]SchemaDefinition{
//   					"name": &SchemaDefinition{
//   						"type": agentcore.SchemaDefinitionType_STRING,
//   						"description": jsii.String("The name to greet"),
//   					},
//   				},
//   				Required: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type AddLambdaTargetOptions struct {
	// The Lambda function to associate with this target.
	// Experimental.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// The tool schema defining the available tools.
	// Experimental.
	ToolSchema ToolSchema `field:"required" json:"toolSchema" yaml:"toolSchema"`
	// Credential providers for authentication.
	// Default: - [GatewayCredentialProvider.iamRole()]
	//
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// Optional description for the gateway target.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
}

