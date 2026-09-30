package awskinesis


// An S3 Tables destination including its table bucket, namespace, table name, and partition spec.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3TableConfigurationProperty := &S3TableConfigurationProperty{
//   	CompressionType: jsii.String("compressionType"),
//   	Namespace: jsii.String("namespace"),
//   	PartitionSpec: &PartitionSpecProperty{
//   		PartitionFields: []interface{}{
//   			&PartitionFieldProperty{
//   				SourceName: jsii.String("sourceName"),
//   				Transform: jsii.String("transform"),
//   			},
//   		},
//   	},
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tableconfiguration.html
//
type CfnChannelPropsMixin_S3TableConfigurationProperty struct {
	// The compression algorithm applied to objects delivered to the S3 Tables destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tableconfiguration.html#cfn-kinesis-channel-s3tableconfiguration-compressiontype
	//
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The name of the S3 Tables namespace that contains the destination table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tableconfiguration.html#cfn-kinesis-channel-s3tableconfiguration-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Partitioning specification for the destination Iceberg table.
	//
	// Follows the S3 Tables / Iceberg PartitionSpec format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tableconfiguration.html#cfn-kinesis-channel-s3tableconfiguration-partitionspec
	//
	PartitionSpec interface{} `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// The ARN of the S3 Tables table bucket for record delivery.
	//
	// Buckets can be cross-account but must be in the same region as the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tableconfiguration.html#cfn-kinesis-channel-s3tableconfiguration-tablebucketarn
	//
	TableBucketArn *string `field:"optional" json:"tableBucketArn" yaml:"tableBucketArn"`
	// The name of the destination S3 Tables table.
	//
	// The table is created for the customer if it does not yet exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3tableconfiguration.html#cfn-kinesis-channel-s3tableconfiguration-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

