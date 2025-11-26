package previewawsivschatmixins


// The DestinationConfiguration property type describes a location where chat logs will be stored.
//
// Each member represents the configuration of one log destination. For logging, you define only one type of destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationConfigurationProperty := &DestinationConfigurationProperty{
//   	CloudWatchLogs: &CloudWatchLogsDestinationConfigurationProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   	},
//   	Firehose: &FirehoseDestinationConfigurationProperty{
//   		DeliveryStreamName: jsii.String("deliveryStreamName"),
//   	},
//   	S3: &S3DestinationConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-loggingconfiguration-destinationconfiguration.html
//
type CfnLoggingConfigurationPropsMixin_DestinationConfigurationProperty struct {
	// An Amazon CloudWatch Logs destination configuration where chat activity will be logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-loggingconfiguration-destinationconfiguration.html#cfn-ivschat-loggingconfiguration-destinationconfiguration-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// An Amazon Kinesis Data Firehose destination configuration where chat activity will be logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-loggingconfiguration-destinationconfiguration.html#cfn-ivschat-loggingconfiguration-destinationconfiguration-firehose
	//
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// An Amazon S3 destination configuration where chat activity will be logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-loggingconfiguration-destinationconfiguration.html#cfn-ivschat-loggingconfiguration-destinationconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

