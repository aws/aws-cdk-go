package awspipes


// The parameters for using a Lambda function as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetLambdaFunctionParametersProperty := &pipeTargetLambdaFunctionParametersProperty{
//   	invocationType: jsii.String("invocationType"),
//   }
//
type CfnPipe_PipeTargetLambdaFunctionParametersProperty struct {
	// Choose from the following options.
	//
	// - `REQUEST_RESPONSE` (default) - Invoke synchronously.
	// - `FIRE_AND_FORGET` - Invoke asynchronously.
	//
	// For more information, see [Invocation types](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html#pipes-invocation) in the *Amazon EventBridge User Guide* .
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
}

