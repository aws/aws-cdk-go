package awsmsk


// Contains information about storage volumes attached to Amazon MSK broker nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnCluster_StorageInfoProperty struct {
	// EBS volume information.
	EbsStorageInfo interface{} `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

