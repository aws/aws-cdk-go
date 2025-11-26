package previewawsfsxmixins


// Used to configure quotas that define how much storage a user or group can use on an FSx for OpenZFS volume.
//
// For more information, see [Volume properties](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/managing-volumes.html#volume-properties) in the FSx for OpenZFS User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userAndGroupQuotasProperty := &UserAndGroupQuotasProperty{
//   	Id: jsii.Number(123),
//   	StorageCapacityQuotaGiB: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-userandgroupquotas.html
//
type CfnFileSystemPropsMixin_UserAndGroupQuotasProperty struct {
	// The ID of the user or group that the quota applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-userandgroupquotas.html#cfn-fsx-filesystem-userandgroupquotas-id
	//
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// The user or group's storage quota, in gibibytes (GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-userandgroupquotas.html#cfn-fsx-filesystem-userandgroupquotas-storagecapacityquotagib
	//
	StorageCapacityQuotaGiB *float64 `field:"optional" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// Specifies whether the quota applies to a user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-userandgroupquotas.html#cfn-fsx-filesystem-userandgroupquotas-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

