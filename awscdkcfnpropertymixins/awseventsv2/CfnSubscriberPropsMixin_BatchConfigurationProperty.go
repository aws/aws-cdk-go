package awseventsv2


// Configuration for batching events into a single delivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   batchConfigurationProperty := &BatchConfigurationProperty{
//   	MaxBatchSize: jsii.Number(123),
//   	MaxBatchWindowInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-batchconfiguration.html
//
type CfnSubscriberPropsMixin_BatchConfigurationProperty struct {
	// The maximum number of events in a single batch delivered to the target.
	//
	// The maximum depends on the target: 500 for Kinesis Data Streams and Amazon Data Firehose, 100 for Lambda, Step Functions, and AWS::EventsV2::EventBus targets, 10 for Amazon SQS, Amazon SNS, and AWS::Events::EventBus targets, and 1 for API Gateway, API destinations, and universal service integration targets. The service rejects a value above the target's maximum. Fewer events may be delivered when the batch window elapses. When omitted, the default is 10 for Lambda and Step Functions targets and the target's maximum for other targets. The resolved value applied by the service is returned on read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-batchconfiguration.html#cfn-eventsv2-subscriber-batchconfiguration-maxbatchsize
	//
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// The maximum time in seconds to wait for a batch to fill before delivering it, 0-300.
	//
	// The default is 0 (no wait). The resolved value applied by the service is returned on read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-batchconfiguration.html#cfn-eventsv2-subscriber-batchconfiguration-maxbatchwindowinseconds
	//
	MaxBatchWindowInSeconds *float64 `field:"optional" json:"maxBatchWindowInSeconds" yaml:"maxBatchWindowInSeconds"`
}

