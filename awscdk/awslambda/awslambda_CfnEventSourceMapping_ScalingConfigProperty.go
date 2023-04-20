package awslambda


// ( Amazon Simple Queue Service only) The scaling configuration for the event source.
//
// For more information, see [Configuring maximum concurrency for Amazon SQS event sources](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConfigProperty := &ScalingConfigProperty{
//   	MaximumConcurrency: jsii.Number(123),
//   }
//
type CfnEventSourceMapping_ScalingConfigProperty struct {
	// Limits the number of concurrent instances that the Amazon SQS event source can invoke.
	MaximumConcurrency *float64 `field:"optional" json:"maximumConcurrency" yaml:"maximumConcurrency"`
}

