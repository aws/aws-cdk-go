package awspipes


// The parameters for using a Step Functions state machine as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetStateMachineParametersProperty := &pipeTargetStateMachineParametersProperty{
//   	invocationType: jsii.String("invocationType"),
//   }
//
type CfnPipe_PipeTargetStateMachineParametersProperty struct {
	// Specify whether to wait for the state machine to finish or not.
	//
	// Choose from the following options.
	//
	// - `REQUEST_RESPONSE` (default) - Invoke synchronously.
	// - `FIRE_AND_FORGET` - Invoke asynchronously.
	//
	// For more information, see [Invocation types](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html#pipes-invocation) in the *Amazon EventBridge User Guide* .
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
}

