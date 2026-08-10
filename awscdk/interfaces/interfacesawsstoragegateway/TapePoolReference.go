package interfacesawsstoragegateway


// A reference to a TapePool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tapePoolReference := &TapePoolReference{
//   	PoolArn: jsii.String("poolArn"),
//   	PoolId: jsii.String("poolId"),
//   }
//
type TapePoolReference struct {
	// The PoolARN of the TapePool resource.
	PoolArn *string `field:"required" json:"poolArn" yaml:"poolArn"`
	// The PoolId of the TapePool resource.
	PoolId *string `field:"required" json:"poolId" yaml:"poolId"`
}

