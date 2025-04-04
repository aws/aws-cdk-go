package awscdkpipessourcesalpha

import (
	"time"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Parameters for the Kinesis source.
//
// Example:
//   var sourceStream stream
//   var targetQueue queue
//
//
//   pipeSource := sources.NewKinesisSource(sourceStream, &KinesisSourceParameters{
//   	StartingPosition: sources.KinesisStartingPosition_LATEST,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: pipeSource,
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   })
//
// Experimental.
type KinesisSourceParameters struct {
	// The maximum number of records to include in each batch.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-batchsize
	//
	// Default: 1.
	//
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Define the target to send dead-letter queue events to.
	//
	// The dead-letter queue stores any events that are not successfully delivered to a Pipes target after all retry attempts are exhausted.
	// You can then resolve the issue that caused the failed invocations and replay the events at a later time.
	// In some cases, such as when access is denied to a resource, events are sent directly to the dead-letter queue and are not retried.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-deadletterconfig
	//
	// Default: - no dead-letter queue or topic.
	//
	// Experimental.
	DeadLetterTarget interface{} `field:"optional" json:"deadLetterTarget" yaml:"deadLetterTarget"`
	// The maximum length of a time to wait for events.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-maximumbatchingwindowinseconds
	//
	// Default: - the events will be handled immediately.
	//
	// Experimental.
	MaximumBatchingWindow awscdk.Duration `field:"optional" json:"maximumBatchingWindow" yaml:"maximumBatchingWindow"`
	// Discard records older than the specified age.
	//
	// The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, EventBridge never discards old records.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-maximumrecordageinseconds
	//
	// Default: -1 - EventBridge won't discard old records.
	//
	// Experimental.
	MaximumRecordAge awscdk.Duration `field:"optional" json:"maximumRecordAge" yaml:"maximumRecordAge"`
	// Discard records after the specified number of retries.
	//
	// The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, EventBridge retries failed records until the record expires in the event source.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-maximumretryattempts
	//
	// Default: -1 - EventBridge will retry failed records until the record expires in the event source.
	//
	// Experimental.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// Define how to handle item process failures.
	//
	// {@link OnPartialBatchItemFailure.AUTOMATIC_BISECT} halves each batch and will retry each half until all the records are processed or there is one failed message left in the batch.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-onpartialbatchitemfailure
	//
	// Default: off - EventBridge will retry the entire batch.
	//
	// Experimental.
	OnPartialBatchItemFailure OnPartialBatchItemFailure `field:"optional" json:"onPartialBatchItemFailure" yaml:"onPartialBatchItemFailure"`
	// The number of batches to process concurrently from each shard.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-parallelizationfactor
	//
	// Default: 1.
	//
	// Experimental.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// The position in a stream from which to start reading.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-startingposition
	//
	// Experimental.
	StartingPosition KinesisStartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in ISO 8601 format.
	//
	// Example:
	//   NewDate(date.uTC(jsii.Number(1969), jsii.Number(10), jsii.Number(20), jsii.Number(0), jsii.Number(0), jsii.Number(0)))
	//
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html#cfn-pipes-pipe-pipesourcekinesisstreamparameters-startingpositiontimestamp
	//
	// Default: - no starting position timestamp.
	//
	// Experimental.
	StartingPositionTimestamp *time.Time `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
}

