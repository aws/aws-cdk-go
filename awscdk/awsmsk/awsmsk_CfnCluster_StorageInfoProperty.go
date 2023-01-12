package awsmsk


// Contains information about storage volumes attached to MSK broker nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageInfoProperty := &storageInfoProperty{
//   	ebsStorageInfo: &eBSStorageInfoProperty{
//   		provisionedThroughput: &provisionedThroughputProperty{
//   			enabled: jsii.Boolean(false),
//   			volumeThroughput: jsii.Number(123),
//   		},
//   		volumeSize: jsii.Number(123),
//   	},
//   }
//
type CfnCluster_StorageInfoProperty struct {
	// EBS volume information.
	EbsStorageInfo interface{} `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

