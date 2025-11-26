package previewawsimagebuildermixins


// The settings for starting an image pipeline execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imagePipelineExecutionSettingsProperty := &ImagePipelineExecutionSettingsProperty{
//   	DeploymentId: jsii.String("deploymentId"),
//   	OnUpdate: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagepipelineexecutionsettings.html
//
type CfnImagePropsMixin_ImagePipelineExecutionSettingsProperty struct {
	// The deployment ID of the pipeline, used to trigger new image pipeline executions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagepipelineexecutionsettings.html#cfn-imagebuilder-image-imagepipelineexecutionsettings-deploymentid
	//
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Whether to trigger the image pipeline when the pipeline is updated.
	//
	// False by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imagepipelineexecutionsettings.html#cfn-imagebuilder-image-imagepipelineexecutionsettings-onupdate
	//
	OnUpdate interface{} `field:"optional" json:"onUpdate" yaml:"onUpdate"`
}

