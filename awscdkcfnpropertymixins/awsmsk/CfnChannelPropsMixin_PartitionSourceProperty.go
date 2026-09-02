package awsmsk


// Partition source configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   partitionSourceProperty := &PartitionSourceProperty{
//   	SourceName: jsii.String("sourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-partitionsource.html
//
type CfnChannelPropsMixin_PartitionSourceProperty struct {
	// Source name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-partitionsource.html#cfn-msk-channel-partitionsource-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
}

