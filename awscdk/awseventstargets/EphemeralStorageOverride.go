package awseventstargets


// Override ephemeral storage for the task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ephemeralStorageOverride := &EphemeralStorageOverride{
//   	SizeInGiB: jsii.Number(123),
//   }
//
type EphemeralStorageOverride struct {
	// The total amount, in GiB, of ephemeral storage to set for the task.
	//
	// The minimum supported value is 20 GiB and the maximum supported value is 200 GiB.
	SizeInGiB *float64 `field:"required" json:"sizeInGiB" yaml:"sizeInGiB"`
}

