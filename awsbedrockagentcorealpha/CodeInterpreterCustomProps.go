package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a CodeInterpreter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var codeInterpreterNetworkConfiguration CodeInterpreterNetworkConfiguration
//   var role Role
//
//   codeInterpreterCustomProps := &CodeInterpreterCustomProps{
//   	CodeInterpreterCustomName: jsii.String("codeInterpreterCustomName"),
//   	Description: jsii.String("description"),
//   	ExecutionRole: role,
//   	NetworkConfiguration: codeInterpreterNetworkConfiguration,
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type CodeInterpreterCustomProps struct {
	// The name of the code interpreter Valid characters are a-z, A-Z, 0-9, _ (underscore) The name must start with a letter and can be up to 48 characters long Pattern: [a-zA-Z][a-zA-Z0-9_]{0,47}.
	// Default: - auto generate.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CodeInterpreterCustomName *string `field:"optional" json:"codeInterpreterCustomName" yaml:"codeInterpreterCustomName"`
	// Optional description for the code interpreter Valid characters are a-z, A-Z, 0-9, _ (underscore), - (hyphen) and spaces The description can have up to 200 characters.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the code interpreter to access AWS services.
	// Default: - A new role will be created.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Network configuration for code interpreter.
	// Default: - PUBLIC network mode.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	NetworkConfiguration CodeInterpreterNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Tags (optional) A list of key:value pairs of tags to apply to this Code Interpreter resource.
	// Default: {} - no tags.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

