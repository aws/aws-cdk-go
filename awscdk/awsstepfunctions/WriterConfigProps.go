package awsstepfunctions


// Interface for Writer Config props.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   writerConfigProps := &WriterConfigProps{
//   	OutputType: awscdk.Aws_stepfunctions.OutputType_JSON,
//   	Transformation: awscdk.*Aws_stepfunctions.Transformation_NONE,
//   }
//
type WriterConfigProps struct {
	// The format of the Output of the child workflow executions.
	OutputType OutputType `field:"required" json:"outputType" yaml:"outputType"`
	// The transformation to be applied to the Output of the Child Workflow executions.
	Transformation Transformation `field:"required" json:"transformation" yaml:"transformation"`
}

