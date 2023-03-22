package awsmsk


// Specifies whether provisioned throughput is turned on and the volume throughput target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedThroughputProperty := &provisionedThroughputProperty{
//   	enabled: jsii.Boolean(false),
//   	volumeThroughput: jsii.Number(123),
//   }
//
type CfnCluster_ProvisionedThroughputProperty struct {
	// Specifies whether provisioned throughput is turned on for the cluster.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The provisioned throughput rate in MiB per second.
	VolumeThroughput *float64 `field:"optional" json:"volumeThroughput" yaml:"volumeThroughput"`
}

