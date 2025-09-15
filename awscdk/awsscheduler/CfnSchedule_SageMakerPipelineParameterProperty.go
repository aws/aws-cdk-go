package awsscheduler


// The name and value pair of a parameter to use to start execution of a SageMaker Model Building Pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerPipelineParameterProperty := &SageMakerPipelineParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-sagemakerpipelineparameter.html
//
type CfnSchedule_SageMakerPipelineParameterProperty struct {
	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-sagemakerpipelineparameter.html#cfn-scheduler-schedule-sagemakerpipelineparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-sagemakerpipelineparameter.html#cfn-scheduler-schedule-sagemakerpipelineparameter-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

