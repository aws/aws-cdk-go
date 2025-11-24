package mixinsawsmedialive


// The current source for one of the pipelines in the multiplex.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexProgramPipelineDetailProperty := &MultiplexProgramPipelineDetailProperty{
//   	ActiveChannelPipeline: jsii.String("activeChannelPipeline"),
//   	PipelineId: jsii.String("pipelineId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampipelinedetail.html
//
type CfnMultiplexprogramPropsMixin_MultiplexProgramPipelineDetailProperty struct {
	// Identifies the channel pipeline that is currently active for the pipeline (identified by PipelineId) in the multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampipelinedetail.html#cfn-medialive-multiplexprogram-multiplexprogrampipelinedetail-activechannelpipeline
	//
	ActiveChannelPipeline *string `field:"optional" json:"activeChannelPipeline" yaml:"activeChannelPipeline"`
	// Identifies a specific pipeline in the multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampipelinedetail.html#cfn-medialive-multiplexprogram-multiplexprogrampipelinedetail-pipelineid
	//
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
}

