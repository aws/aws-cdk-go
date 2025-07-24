package awsschedulertargets


// Properties for a pipeline parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerPipelineParameter := &SageMakerPipelineParameter{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
type SageMakerPipelineParameter struct {
	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	Value *string `field:"required" json:"value" yaml:"value"`
}

