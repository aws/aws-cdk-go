package awsmsk


// Contains information about the EBS storage volumes attached to the broker nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eBSStorageInfoProperty := &EBSStorageInfoProperty{
//   	ProvisionedThroughput: &ProvisionedThroughputProperty{
//   		Enabled: jsii.Boolean(false),
//   		VolumeThroughput: jsii.Number(123),
//   	},
//   	VolumeSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-ebsstorageinfo.html
//
type CfnCluster_EBSStorageInfoProperty struct {
	// EBS volume provisioned throughput information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-ebsstorageinfo.html#cfn-msk-cluster-ebsstorageinfo-provisionedthroughput
	//
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-ebsstorageinfo.html#cfn-msk-cluster-ebsstorageinfo-volumesize
	//
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

