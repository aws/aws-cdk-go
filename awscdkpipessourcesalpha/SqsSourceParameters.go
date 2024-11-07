package awscdkpipessourcesalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Parameters for the SQS source.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//
//
//   pipeSource := sources.NewSqsSource(sourceQueue, &SqsSourceParameters{
//   	BatchSize: jsii.Number(10),
//   	MaximumBatchingWindow: cdk.Duration_Seconds(jsii.Number(10)),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: pipeSource,
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   })
//
// Experimental.
type SqsSourceParameters struct {
	// The maximum number of records to include in each batch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcesqsqueueparameters.html#cfn-pipes-pipe-pipesourcesqsqueueparameters-batchsize
	//
	// Default: 10.
	//
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// The maximum length of a time to wait for events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipesourcesqsqueueparameters.html#cfn-pipes-pipe-pipesourcesqsqueueparameters-maximumbatchingwindowinseconds
	//
	// Default: 1.
	//
	// Experimental.
	MaximumBatchingWindow awscdk.Duration `field:"optional" json:"maximumBatchingWindow" yaml:"maximumBatchingWindow"`
}

