package awsscheduler


// The templated target type for the Amazon Kinesis [`PutRecord`](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecord.html) API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisParametersProperty := &KinesisParametersProperty{
//   	PartitionKey: jsii.String("partitionKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-kinesisparameters.html
//
type CfnSchedule_KinesisParametersProperty struct {
	// Specifies the shard to which EventBridge Scheduler sends the event.
	//
	// For more information, see [Amazon Kinesis Data Streams terminology and concepts](https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html) in the *Amazon Kinesis Streams Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-kinesisparameters.html#cfn-scheduler-schedule-kinesisparameters-partitionkey
	//
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}

