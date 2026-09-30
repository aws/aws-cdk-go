package awskinesis


// Partitioning specification for the destination Iceberg table.
//
// Follows the S3 Tables / Iceberg PartitionSpec format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   partitionSpecProperty := &PartitionSpecProperty{
//   	PartitionFields: []interface{}{
//   		&PartitionFieldProperty{
//   			SourceName: jsii.String("sourceName"),
//   			Transform: jsii.String("transform"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-partitionspec.html
//
type CfnChannelPropsMixin_PartitionSpecProperty struct {
	// List of partition fields that define how records are partitioned when written to the destination table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-partitionspec.html#cfn-kinesis-channel-partitionspec-partitionfields
	//
	PartitionFields interface{} `field:"optional" json:"partitionFields" yaml:"partitionFields"`
}

