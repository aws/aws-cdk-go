package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseProperty := &FirehoseProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	DeliveryStream: jsii.String("deliveryStream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-firehose.html
//
type CfnCluster_FirehoseProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-firehose.html#cfn-msk-cluster-firehose-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-firehose.html#cfn-msk-cluster-firehose-deliverystream
	//
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

