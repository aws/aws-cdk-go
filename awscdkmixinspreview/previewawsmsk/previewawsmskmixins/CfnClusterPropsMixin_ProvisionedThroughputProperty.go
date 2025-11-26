package previewawsmskmixins


// Contains information about provisioned throughput for EBS storage volumes attached to kafka broker nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   provisionedThroughputProperty := &ProvisionedThroughputProperty{
//   	Enabled: jsii.Boolean(false),
//   	VolumeThroughput: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-provisionedthroughput.html
//
type CfnClusterPropsMixin_ProvisionedThroughputProperty struct {
	// Provisioned throughput is on or off.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-provisionedthroughput.html#cfn-msk-cluster-provisionedthroughput-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Throughput value of the EBS volumes for the data drive on each kafka broker node in MiB per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-provisionedthroughput.html#cfn-msk-cluster-provisionedthroughput-volumethroughput
	//
	VolumeThroughput *float64 `field:"optional" json:"volumeThroughput" yaml:"volumeThroughput"`
}

