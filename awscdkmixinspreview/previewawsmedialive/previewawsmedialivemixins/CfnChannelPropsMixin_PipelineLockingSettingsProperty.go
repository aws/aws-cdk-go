package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipelineLockingSettingsProperty := &PipelineLockingSettingsProperty{
//   	PipelineLockingMethod: jsii.String("pipelineLockingMethod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-pipelinelockingsettings.html
//
type CfnChannelPropsMixin_PipelineLockingSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-pipelinelockingsettings.html#cfn-medialive-channel-pipelinelockingsettings-pipelinelockingmethod
	//
	PipelineLockingMethod *string `field:"optional" json:"pipelineLockingMethod" yaml:"pipelineLockingMethod"`
}

