package awseventsv2


// Parameters for writing events to an Amazon Kinesis Data Streams target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kinesisParametersProperty := &KinesisParametersProperty{
//   	ExplicitHashKey: jsii.String("explicitHashKey"),
//   	PartitionKey: jsii.String("partitionKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-kinesisparameters.html
//
type CfnSubscriberPropsMixin_KinesisParametersProperty struct {
	// An explicit hash key that overrides the partition key's shard assignment.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-kinesisparameters.html#cfn-eventsv2-subscriber-kinesisparameters-explicithashkey
	//
	ExplicitHashKey *string `field:"optional" json:"explicitHashKey" yaml:"explicitHashKey"`
	// The partition key that determines which shard each record is written to.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-kinesisparameters.html#cfn-eventsv2-subscriber-kinesisparameters-partitionkey
	//
	PartitionKey *string `field:"optional" json:"partitionKey" yaml:"partitionKey"`
}

