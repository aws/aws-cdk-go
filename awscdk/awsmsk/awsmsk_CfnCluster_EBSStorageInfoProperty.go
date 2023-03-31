package awsmsk


// Contains information about the EBS storage volumes attached to brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eBSStorageInfoProperty := &eBSStorageInfoProperty{
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		enabled: jsii.Boolean(false),
//   		volumeThroughput: jsii.Number(123),
//   	},
//   	volumeSize: jsii.Number(123),
//   }
//
type CfnCluster_EBSStorageInfoProperty struct {
	// Specifies whether provisioned throughput is turned on and the volume throughput target.
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

