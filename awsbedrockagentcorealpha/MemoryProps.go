package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Memory resource.
//
// Example:
//   // Create a custom execution role
//   executionRole := iam.NewRole(this, jsii.String("MemoryExecutionRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("bedrock-agentcore.amazonaws.com")),
//   	ManagedPolicies: []IManagedPolicy{
//   		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonBedrockAgentCoreMemoryBedrockModelInferenceExecutionRolePolicy")),
//   	},
//   })
//
//   // Create memory with custom execution role
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my_memory"),
//   	Description: jsii.String("Memory with custom execution role"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	ExecutionRole: executionRole,
//   })
//
// Experimental.
type MemoryProps struct {
	// The name of the memory Valid characters are a-z, A-Z, 0-9, _ (underscore) The name must start with a letter and can be up to 48 characters long Pattern: [a-zA-Z][a-zA-Z0-9_]{0,47}.
	// Experimental.
	MemoryName *string `field:"required" json:"memoryName" yaml:"memoryName"`
	// Optional description for the memory Valid characters are a-z, A-Z, 0-9, _ (underscore), - (hyphen) and spaces The description can have up to 200 characters.
	// Default: No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the memory to access AWS services when using custom strategies.
	// Default: A new role will be created.
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Short-term memory expiration in days (between 7 and 365).
	//
	// Sets the short-term (raw event) memory retention.
	// Events older than the specified duration will expire and no longer be stored.
	// Default: 90.
	//
	// Experimental.
	ExpirationDuration awscdk.Duration `field:"optional" json:"expirationDuration" yaml:"expirationDuration"`
	// Custom KMS key to use for encryption.
	// Default: Your data is encrypted with a key that AWS owns and manages for you.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// If you need long-term memory for context recall across sessions, you can setup memory extraction strategies to extract the relevant memory from the raw events.
	// Default: No extraction strategies (short term memory only).
	//
	// Experimental.
	MemoryStrategies *[]IMemoryStrategy `field:"optional" json:"memoryStrategies" yaml:"memoryStrategies"`
	// Tags (optional) A list of key:value pairs of tags to apply to this memory resource.
	// Default: {} no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

