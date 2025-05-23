package awselasticache


// Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationDetailsProperty := &DestinationDetailsProperty{
//   	CloudWatchLogsDetails: &CloudWatchLogsDestinationDetailsProperty{
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	KinesisFirehoseDetails: &KinesisFirehoseDestinationDetailsProperty{
//   		DeliveryStream: jsii.String("deliveryStream"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cachecluster-destinationdetails.html
//
type CfnCacheCluster_DestinationDetailsProperty struct {
	// The configuration details of the CloudWatch Logs destination.
	//
	// Note that this field is marked as required but only if CloudWatch Logs was chosen as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cachecluster-destinationdetails.html#cfn-elasticache-cachecluster-destinationdetails-cloudwatchlogsdetails
	//
	CloudWatchLogsDetails interface{} `field:"optional" json:"cloudWatchLogsDetails" yaml:"cloudWatchLogsDetails"`
	// The configuration details of the Kinesis Data Firehose destination.
	//
	// Note that this field is marked as required but only if Kinesis Data Firehose was chosen as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cachecluster-destinationdetails.html#cfn-elasticache-cachecluster-destinationdetails-kinesisfirehosedetails
	//
	KinesisFirehoseDetails interface{} `field:"optional" json:"kinesisFirehoseDetails" yaml:"kinesisFirehoseDetails"`
}

