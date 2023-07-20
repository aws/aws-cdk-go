package awspipes


// The parameters for using a DynamoDB stream as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceDynamoDBStreamParametersProperty := &PipeSourceDynamoDBStreamParametersProperty{
//   	StartingPosition: jsii.String("startingPosition"),
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	DeadLetterConfig: &DeadLetterConfigProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	MaximumRecordAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   	OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   	ParallelizationFactor: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html
//
type CfnPipe_PipeSourceDynamoDBStreamParametersProperty struct {
	// (Streams only) The position in a stream from which to start reading.
	//
	// *Valid values* : `TRIM_HORIZON | LATEST`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html#cfn-pipes-pipe-pipesourcedynamodbstreamparameters-startingposition
	//
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The maximum number of records to include in each batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html#cfn-pipes-pipe-pipesourcedynamodbstreamparameters-batchsize
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Define the target queue to send dead-letter queue events to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html#cfn-pipes-pipe-pipesourcedynamodbstreamparameters-deadletterconfig
	//
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The maximum length of a time to wait for events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html#cfn-pipes-pipe-pipesourcedynamodbstreamparameters-maximumbatchingwindowinseconds
	//
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// (Streams only) Discard records older than the specified age.
	//
	// The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, EventBridge never discards old records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html#cfn-pipes-pipe-pipesourcedynamodbstreamparameters-maximumrecordageinseconds
	//
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// (Streams only) Discard records after the specified number of retries.
	//
	// The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, EventBridge retries failed records until the record expires in the event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html#cfn-pipes-pipe-pipesourcedynamodbstreamparameters-maximumretryattempts
	//
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// (Streams only) Define how to handle item process failures.
	//
	// `AUTOMATIC_BISECT` halves each batch and retry each half until all the records are processed or there is one failed message left in the batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html#cfn-pipes-pipe-pipesourcedynamodbstreamparameters-onpartialbatchitemfailure
	//
	OnPartialBatchItemFailure *string `field:"optional" json:"onPartialBatchItemFailure" yaml:"onPartialBatchItemFailure"`
	// (Streams only) The number of batches to process concurrently from each shard.
	//
	// The default value is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html#cfn-pipes-pipe-pipesourcedynamodbstreamparameters-parallelizationfactor
	//
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
}

