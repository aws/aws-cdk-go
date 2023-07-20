package awsfsx


// An object specifying how much storage users or groups can use on the volume.
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
	// The ID of the user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-userandgroupquotas.html#cfn-fsx-volume-userandgroupquotas-id
	//
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// The amount of storage that the user or group can use in gibibytes (GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-userandgroupquotas.html#cfn-fsx-volume-userandgroupquotas-storagecapacityquotagib
	//
	StorageCapacityQuotaGiB *float64 `field:"required" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// A value that specifies whether the quota applies to a user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-userandgroupquotas.html#cfn-fsx-volume-userandgroupquotas-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

