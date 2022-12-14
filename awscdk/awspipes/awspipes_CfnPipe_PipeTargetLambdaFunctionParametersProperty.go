package awspipes


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
	// `CfnPipe.PipeTargetLambdaFunctionParametersProperty.InvocationType`.
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
}

