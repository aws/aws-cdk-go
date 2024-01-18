package awsfsx


// Configures how much storage users and groups can use on the volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userAndGroupQuotasProperty := &UserAndGroupQuotasProperty{
//   	Id: jsii.Number(123),
//   	StorageCapacityQuotaGiB: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-userandgroupquotas.html
//
type CfnVolume_UserAndGroupQuotasProperty struct {
	// The ID of the user or group that the quota applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-userandgroupquotas.html#cfn-fsx-volume-userandgroupquotas-id
	//
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// The user or group's storage quota, in gibibytes (GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-userandgroupquotas.html#cfn-fsx-volume-userandgroupquotas-storagecapacityquotagib
	//
	StorageCapacityQuotaGiB *float64 `field:"required" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// Specifies whether the quota applies to a user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-userandgroupquotas.html#cfn-fsx-volume-userandgroupquotas-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

