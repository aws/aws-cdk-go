package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ephemeralStorageProperty := &EphemeralStorageProperty{
//   	SizeInGiB: jsii.Number(123),
//   }
//
type CfnJobDefinition_EphemeralStorageProperty struct {
	// `CfnJobDefinition.EphemeralStorageProperty.SizeInGiB`.
	SizeInGiB *float64 `field:"required" json:"sizeInGiB" yaml:"sizeInGiB"`
}

