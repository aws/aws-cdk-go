package awsevents


// These are custom parameters to use when the target is a SageMaker Model Building Pipeline that starts based on EventBridge events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sageMakerPipelineParametersProperty := &SageMakerPipelineParametersProperty{
//   	PipelineParameterList: []interface{}{
//   		&SageMakerPipelineParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRule_SageMakerPipelineParametersProperty struct {
	// List of Parameter names and values for SageMaker Model Building Pipeline execution.
	PipelineParameterList interface{} `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

