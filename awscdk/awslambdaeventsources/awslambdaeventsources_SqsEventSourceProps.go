package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"), &queueProps{
//   	visibilityTimeout: awscdk.Duration.seconds(jsii.Number(30)),
//   	 // default,
//   	receiveMessageWaitTime: awscdk.Duration.seconds(jsii.Number(20)),
//   })
//
//   fn.addEventSource(awscdk.NewSqsEventSource(queue, &sqsEventSourceProps{
//   	batchSize: jsii.Number(10),
//   	 // default
//   	maxBatchingWindow: awscdk.Duration.minutes(jsii.Number(5)),
//   	reportBatchItemFailures: jsii.Boolean(true),
//   }))
//
type SqsEventSourceProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range: Minimum value of 1. Maximum value of 10.
	// If `maxBatchingWindow` is configured, this value can go up to 10,000.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the SQS event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Valid Range: Minimum value of 0 minutes. Maximum value of 5 minutes.
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#services-sqs-batchfailurereporting
	//
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
}

