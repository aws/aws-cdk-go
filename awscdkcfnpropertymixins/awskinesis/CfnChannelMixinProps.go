package awskinesis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnChannelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnChannelMixinProps := &CfnChannelMixinProps{
//   	ChannelName: jsii.String("channelName"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   		KeyId: jsii.String("keyId"),
//   	},
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		CloudWatchLogs: &CloudWatchLogsConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		DataFreshnessInSeconds: jsii.Number(123),
//   		DeadLetterQueueS3Configuration: &DeadLetterQueueS3ConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		StorageConfiguration: &S3StorageConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			CompressionType: jsii.String("compressionType"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   			StorageClass: jsii.String("storageClass"),
//   		},
//   	},
//   	S3TablesDestinationConfiguration: &S3TablesDestinationConfigurationProperty{
//   		DataFreshnessInSeconds: jsii.Number(123),
//   		DeadLetterQueueS3Configuration: &DeadLetterQueueS3ConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		},
//   		S3TablesConfigurationList: []interface{}{
//   			&S3TableConfigurationProperty{
//   				CompressionType: jsii.String("compressionType"),
//   				Namespace: jsii.String("namespace"),
//   				PartitionSpec: &PartitionSpecProperty{
//   					PartitionFields: []interface{}{
//   						&PartitionFieldProperty{
//   							SourceName: jsii.String("sourceName"),
//   							Transform: jsii.String("transform"),
//   						},
//   					},
//   				},
//   				TableBucketArn: jsii.String("tableBucketArn"),
//   				TableName: jsii.String("tableName"),
//   			},
//   		},
//   	},
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   	StreamConfigurationList: []interface{}{
//   		&StreamConfigurationProperty{
//   			RecordConfiguration: &RecordConfigurationProperty{
//   				GsrSchemaArn: jsii.String("gsrSchemaArn"),
//   				RecordFormatType: jsii.String("recordFormatType"),
//   			},
//   			StreamArn: jsii.String("streamArn"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html
//
type CfnChannelMixinProps struct {
	// The name of the channel.
	//
	// The name's uniqueness is scoped per AWS account and region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html#cfn-kinesis-channel-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Server-side encryption configuration for data at rest in the destination.
	//
	// Data delivered to S3 / S3 Tables is encrypted with the same key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html#cfn-kinesis-channel-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Configuration for delivering channel operational logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html#cfn-kinesis-channel-loggingconfiguration
	//
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Configuration for delivery to a vanilla S3 bucket destination.
	//
	// Exactly one of S3DestinationConfiguration and S3TablesDestinationConfiguration must be specified on the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html#cfn-kinesis-channel-s3destinationconfiguration
	//
	S3DestinationConfiguration interface{} `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// Configuration for delivery to S3 Tables destinations.
	//
	// Exactly one of S3DestinationConfiguration and S3TablesDestinationConfiguration must be specified on the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html#cfn-kinesis-channel-s3tablesdestinationconfiguration
	//
	S3TablesDestinationConfiguration interface{} `field:"optional" json:"s3TablesDestinationConfiguration" yaml:"s3TablesDestinationConfiguration"`
	// The ARN of the IAM role that the channel assumes to read from the source stream, deliver records to the destination, and (when enabled) write CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html#cfn-kinesis-channel-serviceexecutionrolearn
	//
	ServiceExecutionRoleArn *string `field:"optional" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// List of stream configurations associated with the channel.
	//
	// v1 supports a single element; the list shape allows future extensibility to fan in from multiple streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html#cfn-kinesis-channel-streamconfigurationlist
	//
	StreamConfigurationList interface{} `field:"optional" json:"streamConfigurationList" yaml:"streamConfigurationList"`
	// An arbitrary set of tags (key-value pairs) to associate with the Kinesis channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-channel.html#cfn-kinesis-channel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

