package awskinesisfirehose


// Represents a single field in a `PartitionSpec` .
//
// Amazon Data Firehose is in preview release and is subject to change.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionFieldProperty := &PartitionFieldProperty{
//   	SourceName: jsii.String("sourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-partitionfield.html
//
type CfnDeliveryStream_PartitionFieldProperty struct {
	// The column name to be configured in partition spec.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-partitionfield.html#cfn-kinesisfirehose-deliverystream-partitionfield-sourcename
	//
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
}

