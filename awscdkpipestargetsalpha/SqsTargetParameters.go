package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// SQS target properties.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//
//
//   pipeTarget := targets.NewSqsTarget(targetQueue, &SqsTargetParameters{
//   	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
//   		"SomeKey": pipes.DynamicInput_fromEventPath(jsii.String("$.body")),
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// Experimental.
type SqsTargetParameters struct {
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: - none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of sent messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetsqsqueueparameters.html#cfn-pipes-pipe-pipetargetsqsqueueparameters-messagededuplicationid
	//
	// Default: - none.
	//
	// Experimental.
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The FIFO message group ID to use as the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetsqsqueueparameters.html#cfn-pipes-pipe-pipetargetsqsqueueparameters-messagegroupid
	//
	// Default: - none.
	//
	// Experimental.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

