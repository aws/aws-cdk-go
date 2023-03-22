package awsfsx


// An object specifying how much storage users or groups can use on the volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userAndGroupQuotasProperty := &userAndGroupQuotasProperty{
//   	id: jsii.Number(123),
//   	storageCapacityQuotaGiB: jsii.Number(123),
//   	type: jsii.String("type"),
//   }
//
type CfnVolume_UserAndGroupQuotasProperty struct {
	// The ID of the user or group.
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// The amount of storage that the user or group can use in gibibytes (GiB).
	StorageCapacityQuotaGiB *float64 `field:"required" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// A value that specifies whether the quota applies to a user or group.
	Type *string `field:"required" json:"type" yaml:"type"`
}

