package awslambda


// A destination for events that failed processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onFailureProperty := &OnFailureProperty{
//   	Destination: jsii.String("destination"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-onfailure.html
//
type CfnEventSourceMapping_OnFailureProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	//
	// To retain records of [asynchronous invocations](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations) , you can configure an Amazon SNS topic, Amazon SQS queue, Lambda function, or Amazon EventBridge event bus as the destination.
	//
	// To retain records of failed invocations from [Kinesis and DynamoDB event sources](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#event-source-mapping-destinations) , you can configure an Amazon SNS topic or Amazon SQS queue as the destination.
	//
	// To retain records of failed invocations from [self-managed Kafka](https://docs.aws.amazon.com/lambda/latest/dg/with-kafka.html#services-smaa-onfailure-destination) or [Amazon MSK](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-onfailure-destination) , you can configure an Amazon SNS topic, Amazon SQS queue, or Amazon S3 bucket as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-onfailure.html#cfn-lambda-eventsourcemapping-onfailure-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

