package awslambda


// A destination for events that failed processing.
//
// For more information, see [Adding a destination](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations) .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-onfailure.html
//
type CfnEventInvokeConfig_OnFailureProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	//
	// To retain records of unsuccessful [asynchronous invocations](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations) , you can configure an Amazon SNS topic, Amazon SQS queue, Amazon S3 bucket, Lambda function, or Amazon EventBridge event bus as the destination.
	//
	// To retain records of failed invocations from [Kinesis](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html) , [DynamoDB](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html) , [self-managed Kafka](https://docs.aws.amazon.com/lambda/latest/dg/with-kafka.html#services-smaa-onfailure-destination) or [Amazon MSK](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-onfailure-destination) , you can configure an Amazon SNS topic, Amazon SQS queue, or Amazon S3 bucket as the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventinvokeconfig-onfailure.html#cfn-lambda-eventinvokeconfig-onfailure-destination
	//
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

