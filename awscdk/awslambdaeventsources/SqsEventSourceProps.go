package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"), &QueueProps{
//   	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(30)),
//   })
//
//   fn.AddEventSource(awscdk.NewSqsEventSource(queue, &SqsEventSourceProps{
//   	BatchSize: jsii.Number(10),
//   	 // default
//   	MaxBatchingWindow: awscdk.Duration_Minutes(jsii.Number(5)),
//   	ReportBatchItemFailures: jsii.Boolean(true),
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
	// Default: 10.
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the SQS event source mapping should be enabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Add filter criteria option.
	// Default: - None.
	//
	Filters *[]*map[string]interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Valid Range: Minimum value of 0 minutes. Maximum value of 5 minutes.
	// Default: - no batching window. The lambda function will be invoked immediately with the records that are available.
	//
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The maximum concurrency setting limits the number of concurrent instances of the function that an Amazon SQS event source can invoke.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency
	//
	// Valid Range: Minimum value of 2. Maximum value of 1000.
	//
	// Default: - No specific limit.
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#services-sqs-batchfailurereporting
	//
	// Default: false.
	//
	ReportBatchItemFailures *bool `field:"optional" json:"reportBatchItemFailures" yaml:"reportBatchItemFailures"`
}

