package awskinesis


// A single partition field consisting of a transform applied to a source column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   partitionFieldProperty := &PartitionFieldProperty{
//   	SourceName: jsii.String("sourceName"),
//   	Transform: jsii.String("transform"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-partitionfield.html
//
type CfnChannelPropsMixin_PartitionFieldProperty struct {
	// The name of the source column on which the transform is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-partitionfield.html#cfn-kinesis-channel-partitionfield-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
	// The partitioning transform applied to the SourceName column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-partitionfield.html#cfn-kinesis-channel-partitionfield-transform
	//
	Transform *string `field:"optional" json:"transform" yaml:"transform"`
}

