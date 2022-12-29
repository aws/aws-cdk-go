package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceDynamoDBStreamParametersProperty := &pipeSourceDynamoDBStreamParametersProperty{
//   	startingPosition: jsii.String("startingPosition"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	deadLetterConfig: &deadLetterConfigProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	maximumRecordAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   	onPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   	parallelizationFactor: jsii.Number(123),
//   }
//
type CfnPipe_PipeSourceDynamoDBStreamParametersProperty struct {
	// `CfnPipe.PipeSourceDynamoDBStreamParametersProperty.StartingPosition`.
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// `CfnPipe.PipeSourceDynamoDBStreamParametersProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnPipe.PipeSourceDynamoDBStreamParametersProperty.DeadLetterConfig`.
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// `CfnPipe.PipeSourceDynamoDBStreamParametersProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// `CfnPipe.PipeSourceDynamoDBStreamParametersProperty.MaximumRecordAgeInSeconds`.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// `CfnPipe.PipeSourceDynamoDBStreamParametersProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// `CfnPipe.PipeSourceDynamoDBStreamParametersProperty.OnPartialBatchItemFailure`.
	OnPartialBatchItemFailure *string `field:"optional" json:"onPartialBatchItemFailure" yaml:"onPartialBatchItemFailure"`
	// `CfnPipe.PipeSourceDynamoDBStreamParametersProperty.ParallelizationFactor`.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
}

