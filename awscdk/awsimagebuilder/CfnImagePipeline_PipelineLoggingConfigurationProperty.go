package awsimagebuilder


// The logging configuration that's defined for pipeline execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineLoggingConfigurationProperty := &PipelineLoggingConfigurationProperty{
//   	ImageLogGroupName: jsii.String("imageLogGroupName"),
//   	PipelineLogGroupName: jsii.String("pipelineLogGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-pipelineloggingconfiguration.html
//
type CfnImagePipeline_PipelineLoggingConfigurationProperty struct {
	// The log group name that Image Builder uses for image creation.
	//
	// If not specified, the log group name defaults to `/aws/imagebuilder/image-name` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-pipelineloggingconfiguration.html#cfn-imagebuilder-imagepipeline-pipelineloggingconfiguration-imageloggroupname
	//
	ImageLogGroupName *string `field:"optional" json:"imageLogGroupName" yaml:"imageLogGroupName"`
	// The log group name that Image Builder uses for the log output during creation of a new pipeline.
	//
	// If not specified, the pipeline log group name defaults to `/aws/imagebuilder/pipeline/pipeline-name` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-pipelineloggingconfiguration.html#cfn-imagebuilder-imagepipeline-pipelineloggingconfiguration-pipelineloggroupname
	//
	PipelineLogGroupName *string `field:"optional" json:"pipelineLogGroupName" yaml:"pipelineLogGroupName"`
}

