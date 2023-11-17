package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnPipe_PipeLogConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-cloudwatchlogslogdestination
	//
	CloudwatchLogsLogDestination interface{} `field:"optional" json:"cloudwatchLogsLogDestination" yaml:"cloudwatchLogsLogDestination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-firehoselogdestination
	//
	FirehoseLogDestination interface{} `field:"optional" json:"firehoseLogDestination" yaml:"firehoseLogDestination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-includeexecutiondata
	//
	IncludeExecutionData *[]*string `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-s3logdestination
	//
	S3LogDestination interface{} `field:"optional" json:"s3LogDestination" yaml:"s3LogDestination"`
}

