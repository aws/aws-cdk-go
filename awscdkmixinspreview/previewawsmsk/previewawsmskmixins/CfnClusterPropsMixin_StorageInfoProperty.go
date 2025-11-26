package previewawsmskmixins


// Contains information about storage volumes attached to Amazon MSK broker nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageInfoProperty := &StorageInfoProperty{
//   	EbsStorageInfo: &EBSStorageInfoProperty{
//   		ProvisionedThroughput: &ProvisionedThroughputProperty{
//   			Enabled: jsii.Boolean(false),
//   			VolumeThroughput: jsii.Number(123),
//   		},
//   		VolumeSize: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-storageinfo.html
//
type CfnClusterPropsMixin_StorageInfoProperty struct {
	// EBS volume information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-storageinfo.html#cfn-msk-cluster-storageinfo-ebsstorageinfo
	//
	EbsStorageInfo interface{} `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

