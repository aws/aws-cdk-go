package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Memory resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key Key
//   var memoryStrategy IMemoryStrategy
//   var role Role
//
//   memoryProps := &MemoryProps{
//   	Description: jsii.String("description"),
//   	ExecutionRole: role,
//   	ExpirationDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   	KmsKey: key,
//   	MemoryName: jsii.String("memoryName"),
//   	MemoryStrategies: []IMemoryStrategy{
//   		memoryStrategy,
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type MemoryProps struct {
	// Optional description for the memory Valid characters are a-z, A-Z, 0-9, _ (underscore), - (hyphen) and spaces The description can have up to 200 characters.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the memory to access AWS services when using custom strategies.
	// Default: - A new role will be created.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Short-term memory expiration in days (between 7 and 365).
	//
	// Sets the short-term (raw event) memory retention.
	// Events older than the specified duration will expire and no longer be stored.
	// Default: - 90 days.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ExpirationDuration awscdk.Duration `field:"optional" json:"expirationDuration" yaml:"expirationDuration"`
	// Custom KMS key to use for encryption.
	// Default: - Your data is encrypted with a key that AWS owns and manages for you.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The name of the memory Valid characters are a-z, A-Z, 0-9, _ (underscore) The name must start with a letter and can be up to 48 characters long Pattern: [a-zA-Z][a-zA-Z0-9_]{0,47}.
	// Default: - auto generate.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MemoryName *string `field:"optional" json:"memoryName" yaml:"memoryName"`
	// If you need long-term memory for context recall across sessions, you can setup memory extraction strategies to extract the relevant memory from the raw events.
	// Default: - No extraction strategies (short term memory only).
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MemoryStrategies *[]IMemoryStrategy `field:"optional" json:"memoryStrategies" yaml:"memoryStrategies"`
	// Tags (optional) A list of key:value pairs of tags to apply to this memory resource.
	// Default: - no tags.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

