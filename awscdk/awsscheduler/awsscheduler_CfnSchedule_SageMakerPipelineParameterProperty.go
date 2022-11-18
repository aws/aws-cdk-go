package awsscheduler


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
	// `CfnSchedule.SageMakerPipelineParameterProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnSchedule.SageMakerPipelineParameterProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

