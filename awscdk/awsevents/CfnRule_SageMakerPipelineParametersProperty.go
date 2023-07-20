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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sagemakerpipelineparameters.html
//
type CfnRule_SageMakerPipelineParametersProperty struct {
	// List of Parameter names and values for SageMaker Model Building Pipeline execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sagemakerpipelineparameters.html#cfn-events-rule-sagemakerpipelineparameters-pipelineparameterlist
	//
	PipelineParameterList interface{} `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

