package awskinesis


// Configuration for delivery to S3 Tables destinations.
//
// Exactly one of S3DestinationConfiguration and S3TablesDestinationConfiguration must be specified on the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3TablesDestinationConfigurationProperty := &S3TablesDestinationConfigurationProperty{
//   	DataFreshnessInSeconds: jsii.Number(123),
//   	DeadLetterQueueS3Configuration: &DeadLetterQueueS3ConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	},
//   	S3TablesConfigurationList: []interface{}{
//   		&S3TableConfigurationProperty{
//   			CompressionType: jsii.String("compressionType"),
//   			Namespace: jsii.String("namespace"),
//   			PartitionSpec: &PartitionSpecProperty{
//   				PartitionFields: []interface{}{
//   					&PartitionFieldProperty{
//   						SourceName: jsii.String("sourceName"),
//   						Transform: jsii.String("transform"),
//   					},
//   				},
//   			},
//   			TableBucketArn: jsii.String("tableBucketArn"),
//   			TableName: jsii.String("tableName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tablesdestinationconfiguration.html
//
type CfnChannelPropsMixin_S3TablesDestinationConfigurationProperty struct {
	// The maximum time in seconds the channel buffers records before delivery if the minimum target file size is not reached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tablesdestinationconfiguration.html#cfn-kinesis-channel-s3tablesdestinationconfiguration-datafreshnessinseconds
	//
	// Default: - 300.
	//
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
	// Configuration of the S3 bucket used to capture records that cannot be delivered to the primary destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tablesdestinationconfiguration.html#cfn-kinesis-channel-s3tablesdestinationconfiguration-deadletterqueues3configuration
	//
	DeadLetterQueueS3Configuration interface{} `field:"optional" json:"deadLetterQueueS3Configuration" yaml:"deadLetterQueueS3Configuration"`
	// The list of S3 Tables destinations.
	//
	// v1 supports a single element; the list shape allows future extensibility to fan out to multiple tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tablesdestinationconfiguration.html#cfn-kinesis-channel-s3tablesdestinationconfiguration-s3tablesconfigurationlist
	//
	S3TablesConfigurationList interface{} `field:"optional" json:"s3TablesConfigurationList" yaml:"s3TablesConfigurationList"`
}

