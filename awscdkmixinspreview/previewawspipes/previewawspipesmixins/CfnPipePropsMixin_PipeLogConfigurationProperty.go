package previewawspipesmixins


// Represents the configuration settings for the logs to which this pipe should report events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipeLogConfigurationProperty := &PipeLogConfigurationProperty{
//   	CloudwatchLogsLogDestination: &CloudwatchLogsLogDestinationProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   	FirehoseLogDestination: &FirehoseLogDestinationProperty{
//   		DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	},
//   	IncludeExecutionData: []*string{
//   		jsii.String("includeExecutionData"),
//   	},
//   	Level: jsii.String("level"),
//   	S3LogDestination: &S3LogDestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketOwner: jsii.String("bucketOwner"),
//   		OutputFormat: jsii.String("outputFormat"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html
//
type CfnPipePropsMixin_PipeLogConfigurationProperty struct {
	// The logging configuration settings for the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-cloudwatchlogslogdestination
	//
	CloudwatchLogsLogDestination interface{} `field:"optional" json:"cloudwatchLogsLogDestination" yaml:"cloudwatchLogsLogDestination"`
	// The Amazon Data Firehose logging configuration settings for the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-firehoselogdestination
	//
	FirehoseLogDestination interface{} `field:"optional" json:"firehoseLogDestination" yaml:"firehoseLogDestination"`
	// Whether the execution data (specifically, the `payload` , `awsRequest` , and `awsResponse` fields) is included in the log messages for this pipe.
	//
	// This applies to all log destinations for the pipe.
	//
	// For more information, see [Including execution data in logs](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-logs.html#eb-pipes-logs-execution-data) in the *Amazon EventBridge User Guide* .
	//
	// *Allowed values:* `ALL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-includeexecutiondata
	//
	IncludeExecutionData *[]*string `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// The level of logging detail to include.
	//
	// This applies to all log destinations for the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
	// The Amazon S3 logging configuration settings for the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-s3logdestination
	//
	S3LogDestination interface{} `field:"optional" json:"s3LogDestination" yaml:"s3LogDestination"`
}

