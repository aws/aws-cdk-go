package awspipes


// The parameters for using a SageMaker pipeline as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetSageMakerPipelineParametersProperty := &PipeTargetSageMakerPipelineParametersProperty{
//   	PipelineParameterList: []interface{}{
//   		&SageMakerPipelineParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipe_PipeTargetSageMakerPipelineParametersProperty struct {
	// List of Parameter names and values for SageMaker Model Building Pipeline execution.
	PipelineParameterList interface{} `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

