package awsfsx


// The configuration for how much storage a user or group can use on the volume.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-userandgroupquotas.html
//
type CfnFileSystem_UserAndGroupQuotasProperty struct {
	// The ID of the user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-userandgroupquotas.html#cfn-fsx-filesystem-userandgroupquotas-id
	//
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// The amount of storage that the user or group can use in gibibytes (GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-userandgroupquotas.html#cfn-fsx-filesystem-userandgroupquotas-storagecapacityquotagib
	//
	StorageCapacityQuotaGiB *float64 `field:"optional" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// A value that specifies whether the quota applies to a user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-userandgroupquotas.html#cfn-fsx-filesystem-userandgroupquotas-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

