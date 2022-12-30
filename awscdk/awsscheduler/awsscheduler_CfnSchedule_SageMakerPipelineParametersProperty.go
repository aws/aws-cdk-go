package awsscheduler


// The templated target type for the Amazon SageMaker [`StartPipelineExecution`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StartPipelineExecution.html) API operation.
//
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
	// List of parameter names and values to use when executing the SageMaker Model Building Pipeline.
	PipelineParameterList interface{} `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

