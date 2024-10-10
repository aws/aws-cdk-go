package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// Parameters for the SfnStateMachine target.
//
// Example:
//   var sourceQueue queue
//   var targetStateMachine iStateMachine
//
//
//   pipeTarget := targets.NewSfnStateMachine(targetStateMachine, &SfnStateMachineParameters{
//   	InvocationType: targets.StateMachineInvocationType_FIRE_AND_FORGET,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// Experimental.
type SfnStateMachineParameters struct {
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// Specify whether to invoke the State Machine synchronously (`REQUEST_RESPONSE`) or asynchronously (`FIRE_AND_FORGET`).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetsqsqueueparameters.html#cfn-pipes-pipe-pipetargetsqsqueueparameters-messagededuplicationid
	//
	// Default: StateMachineInvocationType.FIRE_AND_FORGET
	//
	// Experimental.
	InvocationType StateMachineInvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
}

