package previewawspipesmixins


// The parameters for using a Lambda function as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipeTargetLambdaFunctionParametersProperty := &PipeTargetLambdaFunctionParametersProperty{
//   	InvocationType: jsii.String("invocationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetlambdafunctionparameters.html
//
type CfnPipePropsMixin_PipeTargetLambdaFunctionParametersProperty struct {
	// Specify whether to invoke the function synchronously or asynchronously.
	//
	// - `REQUEST_RESPONSE` (default) - Invoke synchronously. This corresponds to the `RequestResponse` option in the `InvocationType` parameter for the Lambda [Invoke](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax) API.
	// - `FIRE_AND_FORGET` - Invoke asynchronously. This corresponds to the `Event` option in the `InvocationType` parameter for the Lambda [Invoke](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax) API.
	//
	// For more information, see [Invocation types](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html#pipes-invocation) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetlambdafunctionparameters.html#cfn-pipes-pipe-pipetargetlambdafunctionparameters-invocationtype
	//
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
}

