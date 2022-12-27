package awsscheduler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerPipelineParametersProperty := &sageMakerPipelineParametersProperty{
//   	pipelineParameterList: []interface{}{
//   		&sageMakerPipelineParameterProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSchedule_SageMakerPipelineParametersProperty struct {
	// `CfnSchedule.SageMakerPipelineParametersProperty.PipelineParameterList`.
	PipelineParameterList interface{} `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

