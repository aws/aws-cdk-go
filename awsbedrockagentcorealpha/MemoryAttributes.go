package awsbedrockagentcorealpha


// Attributes for specifying an imported Memory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   memoryAttributes := &MemoryAttributes{
//   	MemoryArn: jsii.String("memoryArn"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Status: jsii.String("status"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type MemoryAttributes struct {
	// The ARN of the memory.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MemoryArn *string `field:"required" json:"memoryArn" yaml:"memoryArn"`
	// The ARN of the IAM role associated to the memory.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The created timestamp of the memory.
	// Default: undefined - No created timestamp is provided.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Optional KMS encryption key associated with this memory.
	// Default: undefined - An AWS managed key is used.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The status of the memory.
	// Default: undefined - No status is provided.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// When this memory was last updated.
	// Default: undefined - No last updated timestamp is provided.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

