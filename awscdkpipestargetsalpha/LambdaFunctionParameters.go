package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// Parameters for the LambdaFunction target.
//
// Example:
//   var sourceQueue queue
//   var targetFunction iFunction
//
//
//   pipeTarget := targets.NewLambdaFunction(targetFunction, &LambdaFunctionParameters{
//   	InvocationType: targets.LambdaFunctionInvocationType_FIRE_AND_FORGET,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: NewSomeSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// Experimental.
type LambdaFunctionParameters struct {
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// Specify whether to invoke the Lambda Function synchronously (`REQUEST_RESPONSE`) or asynchronously (`FIRE_AND_FORGET`).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetlambdafunctionparameters.html
	//
	// Default: LambdaFunctionInvocationType.REQUEST_RESPONSE
	//
	// Experimental.
	InvocationType LambdaFunctionInvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
}

