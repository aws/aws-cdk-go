package awsimagebuilder


// Contains settings for starting an image pipeline execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imagePipelineExecutionSettingsProperty := &ImagePipelineExecutionSettingsProperty{
//   	DeploymentId: jsii.String("deploymentId"),
//   	OnUpdate: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagepipelineexecutionsettings.html
//
type CfnImage_ImagePipelineExecutionSettingsProperty struct {
	// The deployment identifier of the pipeline, utilized to initiate new image pipeline executions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagepipelineexecutionsettings.html#cfn-imagebuilder-image-imagepipelineexecutionsettings-deploymentid
	//
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Defines whether the pipeline should be executed upon pipeline updates.
	//
	// False by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagepipelineexecutionsettings.html#cfn-imagebuilder-image-imagepipelineexecutionsettings-onupdate
	//
	OnUpdate interface{} `field:"optional" json:"onUpdate" yaml:"onUpdate"`
}

