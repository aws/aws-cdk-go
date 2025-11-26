package previewawselasticachemixins


// The configuration details of the Kinesis Data Firehose destination.
//
// Note that this field is marked as required but only if Kinesis Data Firehose was chosen as the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisFirehoseDestinationDetailsProperty := &KinesisFirehoseDestinationDetailsProperty{
//   	DeliveryStream: jsii.String("deliveryStream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cachecluster-kinesisfirehosedestinationdetails.html
//
type CfnCacheClusterPropsMixin_KinesisFirehoseDestinationDetailsProperty struct {
	// The name of the Kinesis Data Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cachecluster-kinesisfirehosedestinationdetails.html#cfn-elasticache-cachecluster-kinesisfirehosedestinationdetails-deliverystream
	//
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

