package awsevents


// This object enables you to specify a JSON path to extract from the event and use as the partition key for the Amazon Kinesis data stream, so that you can control the shard to which the event goes.
//
// If you do not include this parameter, the default is to use the `eventId` as the partition key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisParametersProperty := &KinesisParametersProperty{
//   	PartitionKeyPath: jsii.String("partitionKeyPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-kinesisparameters.html
//
type CfnRule_KinesisParametersProperty struct {
	// The JSON path to be extracted from the event and used as the partition key.
	//
	// For more information, see [Amazon Kinesis Streams Key Concepts](https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html#partition-key) in the *Amazon Kinesis Streams Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-kinesisparameters.html#cfn-events-rule-kinesisparameters-partitionkeypath
	//
	PartitionKeyPath *string `field:"required" json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

