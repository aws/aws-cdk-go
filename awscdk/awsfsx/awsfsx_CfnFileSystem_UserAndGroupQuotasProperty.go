package awsfsx


// The configuration for how much storage a user or group can use on the volume.
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
type CfnFileSystem_UserAndGroupQuotasProperty struct {
	// The ID of the user or group.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// The amount of storage that the user or group can use in gibibytes (GiB).
	StorageCapacityQuotaGiB *float64 `field:"optional" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// A value that specifies whether the quota applies to a user or group.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

