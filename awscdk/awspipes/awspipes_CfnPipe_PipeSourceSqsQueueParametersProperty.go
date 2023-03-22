package awspipes


// The parameters for using a Amazon SQS stream as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceSqsQueueParametersProperty := &pipeSourceSqsQueueParametersProperty{
//   	batchSize: jsii.Number(123),
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   }
//
type CfnPipe_PipeSourceSqsQueueParametersProperty struct {
	// The maximum number of records to include in each batch.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The maximum length of a time to wait for events.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
}

