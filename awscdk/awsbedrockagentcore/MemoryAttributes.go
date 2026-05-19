package awsbedrockagentcore


// Attributes for specifying an imported Memory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type MemoryAttributes struct {
	// The ARN of the memory.
	MemoryArn *string `field:"required" json:"memoryArn" yaml:"memoryArn"`
	// The ARN of the IAM role associated to the memory.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The created timestamp of the memory.
	// Default: undefined - No created timestamp is provided.
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Optional KMS encryption key associated with this memory.
	// Default: undefined - An AWS managed key is used.
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The status of the memory.
	// Default: undefined - No status is provided.
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// When this memory was last updated.
	// Default: undefined - No last updated timestamp is provided.
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

