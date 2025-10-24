package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a CodeInterpreter resource.
//
// Example:
//   // Create a custom execution role
//   executionRole := iam.NewRole(this, jsii.String("CodeInterpreterExecutionRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("bedrock-agentcore.amazonaws.com")),
//   })
//
//   // Create code interpreter with custom execution role
//   codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
//   	CodeInterpreterCustomName: jsii.String("my_code_interpreter"),
//   	Description: jsii.String("Code interpreter with custom execution role"),
//   	NetworkConfiguration: agentcore.CodeInterpreterNetworkConfiguration_UsingPublicNetwork(),
//   	ExecutionRole: executionRole,
//   })
//
// Experimental.
type CodeInterpreterCustomProps struct {
	// The name of the code interpreter Valid characters are a-z, A-Z, 0-9, _ (underscore) The name must start with a letter and can be up to 48 characters long Pattern: [a-zA-Z][a-zA-Z0-9_]{0,47}.
	// Experimental.
	CodeInterpreterCustomName *string `field:"required" json:"codeInterpreterCustomName" yaml:"codeInterpreterCustomName"`
	// Optional description for the code interpreter Valid characters are a-z, A-Z, 0-9, _ (underscore), - (hyphen) and spaces The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the code interpreter to access AWS services.
	// Default: - A new role will be created.
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Network configuration for code interpreter.
	// Default: - PUBLIC network mode.
	//
	// Experimental.
	NetworkConfiguration CodeInterpreterNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Tags (optional) A list of key:value pairs of tags to apply to this Code Interpreter resource.
	// Default: {} - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

