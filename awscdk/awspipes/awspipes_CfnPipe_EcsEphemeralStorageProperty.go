package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsEphemeralStorageProperty := &ecsEphemeralStorageProperty{
//   	sizeInGiB: jsii.Number(123),
//   }
//
type CfnPipe_EcsEphemeralStorageProperty struct {
	// `CfnPipe.EcsEphemeralStorageProperty.SizeInGiB`.
	SizeInGiB *float64 `field:"required" json:"sizeInGiB" yaml:"sizeInGiB"`
}

