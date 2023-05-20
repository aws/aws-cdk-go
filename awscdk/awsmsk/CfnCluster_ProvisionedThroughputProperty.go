package awsmsk


// Contains information about provisioned throughput for EBS storage volumes attached to kafka broker nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedThroughputProperty := &ProvisionedThroughputProperty{
//   	Enabled: jsii.Boolean(false),
//   	VolumeThroughput: jsii.Number(123),
//   }
//
type CfnCluster_ProvisionedThroughputProperty struct {
	// Provisioned throughput is enabled or not.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Throughput value of the EBS volumes for the data drive on each kafka broker node in MiB per second.
	VolumeThroughput *float64 `field:"optional" json:"volumeThroughput" yaml:"volumeThroughput"`
}

