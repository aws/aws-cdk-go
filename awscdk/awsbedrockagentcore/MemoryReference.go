package awsbedrockagentcore


// A reference to a Memory resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryReference := &MemoryReference{
//   	MemoryArn: jsii.String("memoryArn"),
//   }
//
type MemoryReference struct {
	// The MemoryArn of the Memory resource.
	MemoryArn *string `field:"required" json:"memoryArn" yaml:"memoryArn"`
}

