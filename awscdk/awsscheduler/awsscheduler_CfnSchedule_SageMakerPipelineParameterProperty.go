package awsscheduler


// The name and value pair of a parameter to use to start execution of a SageMaker Model Building Pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerPipelineParameterProperty := &sageMakerPipelineParameterProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnSchedule_SageMakerPipelineParameterProperty struct {
	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	Value *string `field:"required" json:"value" yaml:"value"`
}

