package awspipes


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
	// `CfnPipe.PipeSourceSqsQueueParametersProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnPipe.PipeSourceSqsQueueParametersProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
}

