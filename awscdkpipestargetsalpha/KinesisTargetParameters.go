package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// Kinesis target properties.
//
// Example:
//   var sourceQueue queue
//   var targetStream stream
//
//
//   streamTarget := targets.NewKinesisTarget(targetStream, &KinesisTargetParameters{
//   	PartitionKey: jsii.String("pk"),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: streamTarget,
//   })
//
// Experimental.
type KinesisTargetParameters struct {
	// Determines which shard in the stream the data record is assigned to.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetkinesisstreamparameters.html#cfn-pipes-pipe-pipetargetkinesisstreamparameters-partitionkey
	//
	// Experimental.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: - none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
}

