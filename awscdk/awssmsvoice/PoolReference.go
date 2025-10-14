package awssmsvoice


// A reference to a Pool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   poolReference := &PoolReference{
//   	PoolArn: jsii.String("poolArn"),
//   	PoolId: jsii.String("poolId"),
//   }
//
type PoolReference struct {
	// The ARN of the Pool resource.
	PoolArn *string `field:"required" json:"poolArn" yaml:"poolArn"`
	// The PoolId of the Pool resource.
	PoolId *string `field:"required" json:"poolId" yaml:"poolId"`
}

