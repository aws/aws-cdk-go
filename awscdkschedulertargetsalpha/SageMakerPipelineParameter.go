package awscdkschedulertargetsalpha


// Properties for a pipeline parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import scheduler_targets_alpha "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha"
//
//   sageMakerPipelineParameter := &SageMakerPipelineParameter{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// Deprecated.
type SageMakerPipelineParameter struct {
	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	// Deprecated.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	// Deprecated.
	Value *string `field:"required" json:"value" yaml:"value"`
}

