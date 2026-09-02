package awsmsk


// Partition specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionSpecProperty := &PartitionSpecProperty{
//   	PartitionStrategy: jsii.String("partitionStrategy"),
//
//   	// the properties below are optional
//   	SourceList: []interface{}{
//   		&PartitionSourceProperty{
//   			SourceName: jsii.String("sourceName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-partitionspec.html
//
type CfnChannel_PartitionSpecProperty struct {
	// Partition strategy for MSK channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-partitionspec.html#cfn-msk-channel-partitionspec-partitionstrategy
	//
	PartitionStrategy *string `field:"required" json:"partitionStrategy" yaml:"partitionStrategy"`
	// Source list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-partitionspec.html#cfn-msk-channel-partitionspec-sourcelist
	//
	SourceList interface{} `field:"optional" json:"sourceList" yaml:"sourceList"`
}

